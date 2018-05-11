package main

import (
	"fmt"
	"io"
)

type M interface {
	io.Writer
}

type S struct{}

func (s S) text() {
	fmt.Println("text")
}

func (s S) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return 10, nil
}
func t(m M) {
	fmt.Println(m)
	m.Write([]byte("kalasa"))
}

func main() {
	s := S{}
	s.Write([]byte("can you help me"))
	t(s)

}
