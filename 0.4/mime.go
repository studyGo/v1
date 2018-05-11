package main

import (
	"fmt"
	"mime"
)

func main() {
	fmt.Println(mime.AddExtensionType(".ws", "application/socket"))
	fmt.Println(mime.TypeByExtension(".ws"))
}
