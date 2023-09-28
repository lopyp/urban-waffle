package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignals(conn net.Conn) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	fmt.Println("Received interrupt signal, closing connection...")
	conn.Close()
	os.Exit(0)
}

func readFromServer(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Connection closed by server.")
			os.Exit(0)
		}
		os.Stdout.Write(buffer[:n])
	}
}

func main() {
	var timeoutStr string
	flag.StringVar(&timeoutStr, "timeout", "10s", "Timeout for connection")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Usage: go-telnet --timeout=10s host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		fmt.Printf("Invalid timeout: %v\n", err)
		os.Exit(1)
	}

	conn, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		fmt.Printf("Error connecting to %s:%s - %v\n", host, port, err)
		os.Exit(1)
	}
	defer conn.Close()

	go handleSignals(conn)
	go readFromServer(conn)

	for {
		data := make([]byte, 1024)
		n, err := os.Stdin.Read(data)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			break
		}
		_, err = conn.Write(data[:n])
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			break
		}
	}
}
