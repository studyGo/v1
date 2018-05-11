package main

import (
    "fmt"
    "net"
)

type so struct {
    read net.Conn
    write net.Conn
}

type Server string
type ClientTable map[net.Conn]*so
type channel chan string

func main () {
    server := ":8080"
    tcpAddr, _ := net.ResolveTCPAddr("tcp", server)
    listen, _ := net.ListenTCP( "tcp", tcpAddr)
    clients := make(ClientTable, 8)
    channel := make(channel, 128)
    go func() {
        for {
            str := <-channel
            for _, client := range clients {
                client.write.Write([]byte(str))
            }
        }
    }()

    for {
        conn, err := listen.Accept()

        channel <- "new connect\n"

        if err != nil {
            fmt.Println("error")
        }
        socketServer := &so {
            read : conn,
            write : conn,
        }

        clients[conn] = socketServer
        clients[conn].write.Write([]byte("welcome to my server\n"))

        buffer := make([]byte, 128)
        go func() {
            for {
                clients[conn].read.Read(buffer)
                channel <- string(buffer)
            }
        }()
        fmt.Println(buffer)
    }
}
