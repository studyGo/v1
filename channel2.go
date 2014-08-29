package main

import "fmt"
import "time"

func main() {

    channel := make(chan string) 
    go func() {
        channel <- "this is a demo"
        time.Sleep(time.Second)
        channel <- "help"
        time.Sleep(time.Second)
        channel <- "sleep"
    }()

    msg1 := <-channel
    fmt.Println(msg1)
    msg2 := <-channel
    fmt.Println(msg2)
    msg3 := <-channel
    fmt.Println(msg3)


}
