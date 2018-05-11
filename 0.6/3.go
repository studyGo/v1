package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "/adn/sdasd/asdasd/"
	s = strings.TrimRight(s, "/")
	fmt.Println(s)

}
