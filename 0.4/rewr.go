package main

import (
	"fmt"
	"time"
)

type Read struct {
	read chan string
}

type Write struct {
	write chan string
}

func main() {
	read := &Read{read: make(chan string, 1)}
	write := &Write{write: make(chan string, 1)}
	end := make(chan string, 1)

	write.write <- "this is a demo"

	go r(read, write)
	go w(read, write)
	<-end
}

func r(re *Read, wr *Write) {
	for {
		data := <-re.read
		fmt.Printf("read : %s", data)
		wr.write <- data
		time.Sleep(time.Second)
	}
}

func w(re *Read, wr *Write) {
	for {
		data := <-wr.write
		fmt.Printf("write : %s", data)
		re.read <- data
		time.Sleep(time.Second)
	}
}
