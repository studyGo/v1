package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Binary string

func (i Binary) String() string {
	return "s"
}

func (i Binary) Get() string {
	return string(i)
}

func m() Binary {
	return "this is a demo"
}

func main() {
	var s Binary
	s = "can you help me"
	fmt.Println(s.String())
	fmt.Println(m().String())
	/*
		s := []Stringer{&b}
		fmt.Print(s.String())
	*/
}
