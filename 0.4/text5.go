package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {

	f, err := os.Create("ss")
	if err != nil {
		os.Exit(0)
	}
	pprof.WriteHeapProfile(f)
	f.Close()
	c := make(chan bool, 1)
	s := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second * 5)
		c <- true
	}()

	go func() {
		time.Sleep(time.Second)
		s <- false
	}()

	select {
	case <-c:
		fmt.Println("c")
	case <-s:
		fmt.Println("s")
	}

}
