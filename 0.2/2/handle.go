package main

import (
	"fmt"
)

type handler string

func (name handler) Test() {
	fmt.Println(name)
}

func Handler(name string) {
	return handler(name)
}

func main() {
	fmt.Println(Handler("demo"))

}
