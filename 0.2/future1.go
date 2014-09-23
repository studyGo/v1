package main

import "fmt"

type demo struct {
    read chan int
    write chan int
}

func do(d demo) {
    go func() {
        var value int = 100
        for {
            select {
            case value = <-d.read:
            case d.write <- value:
            }
        }
    }()
}

func main (){
    d := demo{make(chan int, 1), make(chan int)}
    do(d)
    s := 0;
    for {
        s = <-d.write
        d.read <- s + 1
        fmt.Println(s)
    }
}
