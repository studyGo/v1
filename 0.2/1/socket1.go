package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	server := ":9001"
	tcpAddr, _ := net.ResolveTCPAddr("tcp", server)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	channel := make(chan string, 1000)

	go insertFile(channel)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error")
			continue
		}
		go handleClient(conn, channel)
	}
}

func insertFile(c chan string) {
	file := "/tmp/demo/logs"
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		f.WriteString(<-c)
	}
}

func handleClient(conn net.Conn, c chan string) {
	request := make([]byte, 1024*256)
	read_len, _ := conn.Read(request)

	if read_len > 0 {
		c <- string(request[0:read_len])
	}
}
