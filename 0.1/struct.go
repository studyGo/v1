package main

import "fmt"

type Person struct {
    name string
    age int
    email string
}

func main() {

    person := Person{"何威风", 30, "demo@qq.com"}

    fmt.Println(person)
}
