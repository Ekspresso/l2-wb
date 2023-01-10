package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать простейший telnet-клиент.

// Примеры вызовов:
// go-telnet --timeout=10s host port
// go-telnet mysite.ru 8080
// go-telnet --timeout=3s 1.1.1.1 123

// Требования:
// Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
// После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
// Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
// При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
// Если сокет закрывается со стороны сервера, программа должна также завершаться.
// При подключении к несуществующему сервер, программа должна завершаться через timeout

func main() {
	// Парсинг флага таймаута
	var timeout int
	flag.IntVar(&timeout, "timeout", 10, "таймаут на подключение")
	flag.Parse()

	fmt.Println("Timeout: ", timeout)
	fmt.Println(flag.Args())
	var wg sync.WaitGroup
	// Проверка передаваемых аргументов
	if len(flag.Args()) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s -timeout host port ", os.Args[0])
		os.Exit(1)
	}
	// Преобразование аргументов в адрес
	host := flag.Args()[0]
	port := flag.Args()[1]
	serv := host + ":" + port
	// Вызов функции для подключения к адресу
	connection(serv, time.Duration(timeout)*time.Second, &wg)
}

// Функция подключения к адресу и отправки и обработки данных
func connection(serv string, timeout time.Duration, wg *sync.WaitGroup) {
	conn, err := net.DialTimeout("tcp", serv, timeout) // открываем TCP-соединение к серверу с таймаутом
	if err != nil {
		fmt.Println("Error connection")
		time.Sleep(timeout)
		os.Exit(0)
	}
	defer conn.Close()
	ctxAll, cancelAll := context.WithCancel(context.Background()) // Контекст для обработки системных сигналов
	wg.Add(1)
	go chanSys(ctxAll, cancelAll, wg) // Запуск функции отслеживания системных сигналов
	fmt.Println("The connection is established")
	fmt.Println("Enter the word")
	// Запуск функции отправки в сокет и обработки ответа сокета
	wg.Add(1)
	go func() {
		defer cancelAll()
		defer wg.Done()
		for {
			select {
			// Прослушивание канала с системными сигналами
			case <-ctxAll.Done():
				fmt.Println("Ending by context sys signal")
				return
			default:
				var word []byte
				scan := bufio.NewScanner(os.Stdin)
				// Считывание командной строки и запись строки в переменную  word
				if scan.Scan() {
					word = scan.Bytes()
				} else { // Если был системный сигнал (EOF), то контекст отменяется
					cancelAll()
					return
				}
				// Если длина строки 0, то ничего не происходит
				if len(word) == 0 {
					continue
				}
				fmt.Println("Your word: ", string(word))
				// Установка таймаута на чтение и запись с сервера
				err := conn.SetDeadline(time.Now().Add(timeout))
				if err != nil {
					fmt.Println("err write deadline: ", err.Error())
					cancelAll()
					continue
				}
				// Отправка данны на сервер
				n, err := conn.Write(word)
				if err != nil && n == 0 {
					fmt.Println("Error write: ", err.Error())
					cancelAll()
					continue
				}
				buf := make([]byte, 1024)
				// Чтение ответа сервера
				n, err = conn.Read(buf)
				if err != nil && n == 0 {
					fmt.Println("Error read: ", err.Error())
					cancelAll()
					continue
				}
				// Вывод ответа в stdOut
				fmt.Println(string(buf[:n]))
			}
		}
	}()
	wg.Wait()
}

// Функция отслеживающая системный канал, в который поступают сигналы из вне.
// Также отслеживает канал контекста.
// Функция реагирует на сигнал завершения работы программы и отменяет контекст,
// что передаётся в другие функции, которые следят за контекстом.
// Если контекст был отменён в другом месте, то завершает работу функции.
func chanSys(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшение счётчика процессов на 1 при завершении работы функции.
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		// Завершение работы функции при отменённом контексте.
		fmt.Println("Ending func chanSys")
		return
	case <-signalCh:
		// Отмена контекста при получении сигнала и завершение работы функции.
		cancel()
		return
	}
}
