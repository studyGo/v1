package main

import (
	"errors"
	"fmt"
)

var ErrMyChar = errors.New("short write")

func main() {
	fmt.Println(ErrMyChar)
}
