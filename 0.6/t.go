package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(1*time.Second > 2*time.Millisecond)
}
