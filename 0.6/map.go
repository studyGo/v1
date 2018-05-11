package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type S struct {
	inject.Injector
	name string
}

func main() {
	s := &S{name: "kalasa", Injector: inject.New()}
	fmt.Println(s.Map(s.name))

}
