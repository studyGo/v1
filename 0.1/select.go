package main

import "time"
import "fmt"

func main() {
    channal1 := make(chan string)
    channal2 := make(chan string)

    go func () {
        time.Sleep(time.Second)
        channal1 <- "channel1"
    }()

    go func () {
        time.Sleep(time.Second * 2)
        channal2 <- "channel2"
    }()

    for {
        select {
            case msg1 := <-channal1:
                fmt.Println(msg1)
            case msg2 := <-channal2:
                fmt.Println(msg2)
            default:
                time.Sleep(time.Second)
        }
    }
}
