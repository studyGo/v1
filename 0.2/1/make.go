package main

import (
    "fmt"
)

func main() {
    msg := "this is a demo"
    switch t := msg.(type) {
        default:
            fmt.Println("123")
    }
}

