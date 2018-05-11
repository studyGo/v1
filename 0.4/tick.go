package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(15 * time.Second)

	select {
	case <-ticker:

		goto echo
	}

echo:
	fmt.Println(1)

	fmt.Println(2)
}
