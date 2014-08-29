package main

import "fmt"

func add () func () int {
    i, j := 1, 1
    return func() int {
        var tmp = i + j
        i, j = j, tmp
        return tmp
    }
}

func main() {
    a := add()
    for i := 0; i < 20; i++ {
        fmt.Println(a())
    }
}
