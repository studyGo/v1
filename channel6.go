package main

import "fmt"

func execute(num int) (sum int) {
    sum = num * num
    return
}

func main() {
    c := make(chan int, 4)
    data := []int{1,2,3,4,5,6,7,8,9}


    for _, v := range data {
        go func(v int) {
            num := execute(v)
            c <- num
        }(v)
    }


    d := 0
    for i := 0; i < len(data); i++ {
        s := <-c
        fmt.Println(s)
        d += s
    }

    fmt.Println(d)
}
