package test

import (
	"fmt"
)

type Test struct{}

var DefaultTest = &Test{}

func Use() {
	DefaultTest.Use()
}

func (t *Test) Use() {
	fmt.Println("test use")
}

func Sleep()

func goFunc(arg interface{}) {
	fmt.Println(arg)

}
