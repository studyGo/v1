package main

import "fmt"

type Demo struct {
    name string
    age int
}

func main () {
    d := new(Demo)
    s := Demo{"demo", 100}

    d.name = "demo"
    fmt.Println(d.name)
}

