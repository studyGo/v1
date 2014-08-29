package main

import "fmt"

type rect struct{
    width, height int
}

func (r *rect) area() int {

    return r.width * r.height
}

func (r *rect) per() int {

    return 2 * (r.width + r.height)
}

func main() {
    r := rect{width:100, height:200}

    fmt.Println("area :", r.area())
    fmt.Println("per :", r.per())

}
