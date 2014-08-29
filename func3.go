package main

import "fmt"

func r() (x, y int) {
    x = 100
    y = 200

    return
}

func main() {
    x, y := r()
    fmt.Println(x)
    fmt.Println(y)
}
