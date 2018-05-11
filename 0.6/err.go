package main

import (
	"errors"
	"fmt"
)

func main() {
	// type error
	fmt.Println(errors.New("this is a demo"))
	//type string
	fmt.Println(errors.New("this is a demo").Error())

	fmt.Println(a())
}

func a() (string, error) {

	return "err", errors.New("this is a demo")
}
