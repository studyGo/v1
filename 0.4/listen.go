package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ipv6()
	ipv4()
}

func ipv6() {
	ln, err := net.Listen("tcp6", "[::1]:0")
	if err != nil {
		log.Fatal(err)
	}
	ln.Close()
	fmt.Println("end")
}

func ipv4() {
	ln, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	ln.Close()
	fmt.Println("end")
}
