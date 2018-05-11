package main

import (
	"fmt"
	"reflect"
)

func tt() {

}

type mm func(s string)
type f struct {
	TT string
}

func main() {
	fmt.Println(reflect.TypeOf(mm(nil)))

	fmt.Println(mm(nil))
	fmt.Println((*mm)(nil))
}
