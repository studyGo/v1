package main

import "fmt"
import "math/rand"
import "time"

func main() {
    channel := make(chan string)
    rand.Seed(time.Now().Unix())
    go func() {
        cnt := rand.Intn(10)

        fmt.Println("send to channel")
        fmt.Println(cnt)
        for i := 0; i < cnt; i++ {
            channel <- fmt.Sprintf("this is a demo %d", i)
        }
        close(channel)
    }()

    var more bool = true
    var msg string

    for more {
        select {
            case msg, more =<-channel:
                if more {
                    fmt.Println(msg)
                } else {
                    fmt.Println("channel is empty")
                }
        }
    }
}
