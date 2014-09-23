package main

import "fmt"
import "time"

func main() {
    timeout := make(chan bool, 1)
    go func() {
        time.Sleep(1 * time.Second)
        timeout <- true
    }()

    select {
        case <-timeout:
            fmt.Println("time out")
    }
}

