package main

import "fmt"

func add(a int, b float32) (x int) {
    x = a + int(b)
    return
}

func main() {

    d, e := fmt.Println(add(1, 3))

    fmt.Println(d, e)

}
