package main

import "fmt"
import "math/rand"
import "time"

func get() chan string{
    channel := make(chan string)
    s := "0123456789abcdef"
    rand.Seed(time.Now().Unix())
    go func() {
        for {
            d := rand.Intn(15)
            channel <- s[d:d+1]
        }
    }()
    return channel
}

func demo() string {
    //eb06a63d-cfe9-4211-ac35-d67c80d44339
    channel := get()
    var msg string
    msg = ""
    for i := 0; i < 36; i++ {
        if i == 8 || i == 13 || i == 18 || i == 23 {
            msg += "-"
            continue
        }
        msg += <-channel
    }
    return msg
}

func main() {
    //start := time.Now().Unix()
    start := time.Now().UnixNano()
    for i := 0; i < 100000; i++ {
        demo()
    }
    fmt.Println(time.Now().UnixNano() - start)
}
