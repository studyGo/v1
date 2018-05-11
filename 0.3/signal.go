package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handle() {
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("sss")
		}
	}()
}

func main() {
	signalChan := make(chan os.Signal, 1)
	signalhubChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(signalhubChan, syscall.SIGHUP)
	handle()
	go func() {
		<-signalhubChan

		fmt.Println("hub")
	}()
	<-signalChan
	fmt.Println("stop")
}
