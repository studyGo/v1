package main

import (
	"fmt"
)

type n struct {
	name string
	data string
}

func main() {
	s := new(n)
	n.name = "tel"
	fmt.Println(s)

}
