package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := `d`
	data := regexp.MustCompile(`\w`)
	pattern = data.ReplaceAllStringFunc(pattern, func(m string) string {
		return "abc"
	})

	fmt.Println(pattern)
}
