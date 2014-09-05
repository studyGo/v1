package main

import "fmt"
import "time"

func main() {
    ticker := time.NewTicker(time.Second)
    // go 的定时器相当准确
    for t := range ticker.C {
        fmt.Println(t)
    }
}

