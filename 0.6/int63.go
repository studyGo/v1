package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Int63n(time.Now().UnixNano()%1000 + 30))

	fmt.Println(rand.Int63n(10000))
}
