package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool, 1)
	a := 100
	s := func(a int) func() {
		return func() {
			fmt.Println("sleep 30")
			ch <- true
		}
	}

	time.AfterFunc(time.Second, s(a))

	for {

		if ok := <-ch; ok {
			go func() {
				time.AfterFunc(time.Second, s(a))
			}()
		}
	}

}
