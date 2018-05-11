package main

import (
	"fmt"
	"net"
	//"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: socket2.go host:port")
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	file := "/tmp/demo/php_9001.log"
	f, err := os.OpenFile(file, os.O_RDWR, 0666)
	s := make([]byte, 1024)
	n, err := f.Read(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("read len ", n)

	conn.Write(s)

	response := make([]byte, 1024)
	result, _ := conn.Read(response)
	if result != 0 {
		fmt.Printf("%s \n", response[0:])
	}
	os.Exit(0)
}
