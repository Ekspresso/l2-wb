package main

// Реализовать утилиту wget с возможностью скачивать сайты целиком.

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	// заполняем запрос
	request := fillRequest()

	// выполняем запрос
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	//отложенное закрытие дескриптора
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(response.Body)

	// проверяем, что статус запроса успешен
	if response.StatusCode != 200 {
		log.Fatalln("Status Code:", response.Status)
	}

	// проверяем что нам приходит html или файл для сохранения
	fileName := ""
	if strings.Contains(response.Header.Get("Content-Type"), "text/html") {
		fileName = "index.html"
	} else {
		fileName = path.Base(request.URL.Path)
	}

	// создаем файл для записи полученных байт
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	// отложенное закрытие дескриптора
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	// копируем байты из тела в файл
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatalln(err)
	}
}

// Заполняем запрос и возвращаем
func fillRequest() *http.Request {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("Error: need 1 argument - URL")
	}
	url := flag.Arg(0)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// запрос будет выполнен с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	request.WithContext(ctx)
	return request
}
