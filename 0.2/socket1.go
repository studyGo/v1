

// server tcp
package main

import (
    "fmt"
    "net"
)

func main() {
server := ":8080"
    tcpAddr, _ := net.ResolveTCPAddr("tcp", server)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    fmt.Println(err)
    for {
        fmt.Println(2)
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("error")
            continue
        }
        fmt.Println(1)
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    request := make([]byte, 128)
    fmt.Println("handle")
    for {
        read_len, _ := conn.Read(request)

        if read_len == 0 {
            break
        } else {
            fmt.Printf("read %s", request[0:])

            conn.Write([]byte("this is a write"))
        }
    }
}

