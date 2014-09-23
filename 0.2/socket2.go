// socket client

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
    conn.Write([]byte("demo"))

    response := make([]byte, 128)
    result, _ := conn.Read(response)
    if result != 0 {
        fmt.Printf("%s \n", response[0:])
    }
    os.Exit(0)
}

