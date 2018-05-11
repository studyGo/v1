package main

import (
	"fmt"
	"reflect"
)

type st struct {
}

func (this *st) Abc() {
	fmt.Println("asd")
}

func main() {
	s := &st{}
	v := reflect.ValueOf(s)
	v.MethodByName("Abc").Call(nil)
}
