package main

import "fmt"

func main() {
    channal := make(chan string)
    go func() {
        channal <-"hello this is a demo"
    }()

    msg := <-channal
    fmt.Println(msg)

}
