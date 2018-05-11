package main

import (
	"fmt"
	json "github.com/bitly/go-simplejson"
)

func main() {
	s := json.New()
	s.Set("m", "this is a demo")
	//s.Set("", "kalasa")
	fmt.Println(s.String())

}
