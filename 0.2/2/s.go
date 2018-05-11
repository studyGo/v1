package main

import (
	"test"
)

func main() {
	test.Use()

	s := &test.Test{}
	s.Use()
	test.Sleep()
}
