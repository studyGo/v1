package main

import "fmt"
import "time"
import "math/rand"

func r() chan int {
    out := make(chan int)
    rand.Seed(time.Now().Unix())
    go func() {
        for {
            out <- rand.Intn(100)
        }
    }()
    go func() {
        for {
            out <- rand.Intn(100)
        }
    }()
    return out
}

func main() {
    rand := r()
    ticker := time.NewTicker(time.Second)
    for t := range ticker.C {
        fmt.Printf("%d", <-rand)
        fmt.Println(t)
    }
}

