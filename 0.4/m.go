package main

import (
	"demo"
	"fmt"
	"unsafe"
)

func main() {
	t := demo.T{Name: "hello"}
	p := (*demo.T)(unsafe.Pointer(&t))
	//p.SetName("hahah")
	//p.kalasa = "this is use unsafe"
	fmt.Println(t, p)

}
