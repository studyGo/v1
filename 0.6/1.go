package main

import (
	"fmt"
)

func t(a ...string) {

	fmt.Println(a)
}

func main() {

	t("1", "2", "3")
}
