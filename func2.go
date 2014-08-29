package main

import "fmt"

func swap(a, b int) (int ,int) {

    return b, a
}

func main() {
    a := 100
    b := 200

    w, s := swap(a, b)
    fmt.Println(w)
    fmt.Println(s)
}
