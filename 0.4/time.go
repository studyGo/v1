package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	fmt.Println(strconv.FormatInt((time.Now().UnixNano()-t)/1000000000, 10))

}
