package main

import "fmt"
import "math"

type shape interface {
	area() float64
	per() float64
}

// react
type rect struct {
	width, height float64
}

func (r *rect) area() float64 {

	return r.width * r.height
}

func (r *rect) per() float64 {

	return 2 * (r.width + r.height)
}

// circle
type circle struct {
	length float64
}

func (c *circle) area() float64 {
	s := math.Pi * c.length * c.length
	return s
}

func (c *circle) per() float64 {

	s := 2 * math.Pi * c.length
	return s
}

func main() {
	r := rect{100, 200}
	//c := circle{20}

	//s := []shape{&r, &c}
	s := []shape{&r}
	fmt.Println(s)

	/*
		for _, sh := range s {
			fmt.Println(sh.area())
			fmt.Println(sh.per())
		}
	*/
}
