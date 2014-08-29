package main

import "fmt"

func demo(x int, y int) int {
    
    return x + y
}

func main() {
    a := 1
    b := 2
    x := demo(a, b)
    fmt.Println(x)
}
