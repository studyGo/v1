package main

import (
	"fmt"
)

type demo struct {
	s int
	w int
}

var s map[int]demo

func main() {
	s = make(map[int]demo)

	s[0] = demo{s: 10, w: 100}
	s[1] = demo{s: 10, w: 100}
	fmt.Println(s)

	for w, v := range s {
		fmt.Println(w, v.s)
	}
}
