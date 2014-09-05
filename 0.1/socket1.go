package main

import "fmt"
import "net"
import "time"

func main() {
    socket, err := net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
        panic ("error")
    }
    defer socket.Close()
    buf := make([]byte, 1000)

    for i := 0; i < 5; i++ {
        msg := fmt.Sprintf("this is a demo %03d", i)
        n, err := socket.Write([]byte (msg))
        if err != nil {
            fmt.Println("error")
            break
        }
        n, err = socket.Read(buf)
        fmt.Println(string (buf[0:n]))
        time.Sleep(time.Second)
    }
}
