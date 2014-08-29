package main

import "fmt"
import "math/rand"

func demo() chan int {
    out := make(chan int)

    go func() {
        out <- rand.Int()
    }()

    return out
}

func main() {

    c := demo()
    fmt.Printf("%d\n", <-c)
}
