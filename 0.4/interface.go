package main

import (
	"fmt"
)

type S struct{}

func (s *S) Check() {
	fmt.Println(1)
}

type W interface {
	Check()
}

func main() {
	fmt.Println((*W)(nil))

}
