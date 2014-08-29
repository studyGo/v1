package main

import "fmt"

func add(nums ...int) int {
    fmt.Print(nums, " ")
    total := 0

    for _, num := range nums {
        total += num
    }

    return total
}


func main() {
    //nums := []int{1,2,3,4}
    m := make(map[int]string)
    m[0] = "demo"
    m[1] = "haha"

    fmt.Println(m[0])
    delete(m, 0)
    //fmt.Println(add(nums...))
}
