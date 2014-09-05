package main

import "fmt"
import "net"
import "io"

const RECV_BUF_LEN = 1024

func main() {
    socket, err := net.Listen("tcp", "0.0.0.0:8080")
    if err != nil {
        panic("error")
    }
    for {
        conn, err := socket.Accept()
        if err != nil {
            panic("accept erro")
        }

        go server(conn)
    }
}

func server(conn net.Conn) {
    buf := make([]byte, RECV_BUF_LEN)
    defer conn.Close()

    for {
        _, err := conn.Read(buf)
        switch err {
            case nil:
                msg := fmt.Sprintf("this is a demo")
                conn.Write([]byte (msg))
            case io.EOF:
                fmt.Println(err)
                return
            default:
                fmt.Println(err)
                return
        }
    }
}

