package main

import (
	"fmt"
	"time"
)

func main() {

	chanell := make(chan bool, 1)
	chanell1 := make(chan int, 1)

	go func() {
		for {
			time.Sleep(time.Second * 3)
			chanell <- true
		}
	}()

	go func() {
		for {
			chanell1 <- sl()
		}
	}()

	for {
		select {
		case <-chanell:
			fmt.Println("timeout")
		case s := <-chanell1:
			fmt.Println(s)
		}
	}
}

func sl() int {
	time.Sleep(time.Second * 2)
	return 100

}
