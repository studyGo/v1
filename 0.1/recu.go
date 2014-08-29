package main

import "fmt"

func feac(n int) int {
    if n == 0 {
        return 1
    }
    return n * feac(n - 1)

}


func main() {

    fmt.Println(feac(10))

}
