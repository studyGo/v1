package main

import (
    "fmt"
    "net"
)

func main () {
    tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:6379")
    if err != nil {
        fmt.Println("error")
        return
    }
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        fmt.Println("error")
        return
    }
    conn.Write([]byte("*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"))
    fmt.Println(1)

    response := make([]byte, 128)
    result, err := conn.Read(response)
    if err != nil {
        fmt.Println("error")
        return
    }
    fmt.Println(response[:result])
    fmt.Println(result)
}

