package main

import (
	"fmt"
	"log"
	"net"
)

var dict = map[string]string{
	"hello": "привет",
	"blue":  "синий",
	"red":   "красный",
	"white": "белый",
}

func main() {
	// Открываем слушающий сокет localhost:8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	// Цикл чтения из сокета
	for {
		fmt.Println("Server wait connection...")
		conn, err := listener.Accept()
		if err != nil {
			conn.Close()
			log.Println(err)
			continue
		}
		// Запуск горутины, которая обрабатывает сокет. Позволяет обрабатывать несколько подключений
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	fmt.Println("Connection is established")
	for {
		n, err := conn.Read(buf)
		fmt.Println("Word reading: ", string(buf[:n]))
		if n == 0 || err != nil {
			fmt.Println("Read error")
			break
		}
		word := string(buf[:n])
		transl, ok := dict[word]
		if !ok {
			transl = "Word not found"
		}
		_, err = conn.Write([]byte(transl))
		if err != nil {
			fmt.Println("Write error")
			break
		}
	}
}
