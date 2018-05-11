package main

import (
	"fmt"
)

func main() {
	var s interface{}
	s = "demo"

	t, w := s.(string)
	fmt.Println(t, w)

}
