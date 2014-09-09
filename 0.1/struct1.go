package main

import "fmt"

type rects struct{
    width, height int
}

type rect struct{
    rects
    demo int
}

func (r *rect) area() int {

    return r.rects.width * r.rects.height
}

func (r *rect) per() int {

    return 2 * (r.rects.width + r.rects.height)
}
func main() {
    r := rect{rects{100, 200}, 400}
    //r := rect(rects{"width":100})

    fmt.Println("area :", r.area())
    fmt.Println("per :", r.per())

}
