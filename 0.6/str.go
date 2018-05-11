package main

import (
	"fmt"
	"strings"
)

func main() {
	s := make([]byte, 20)
	str := strings.NewReader("this is a demo")
	fmt.Println(str.Len())
	fmt.Println(str.Read(s))
	fmt.Println(string(s[0:]))
	//str.Reset("this is a kalasa")
	fmt.Println(str.Len())
	fmt.Println(2 << 1)

}
