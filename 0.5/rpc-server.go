package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type S int

func (s *S) Multiply(arg string, reply *int) error {
	*reply = 100
	return nil
}

func (s *S) Get(arg string, reply *string) error {
	*reply = "tjhis  s a da"
	return nil
}

func (s *S) Put(arg string, reply *string) error {
	*reply = "Put"
	return nil
}

func (s *S) Heart(arg string, reply *string) error {
	*reply = "Heart"
	return nil
}

func main() {
	s := new(S)
	c := make(chan bool, 1)
	rpc.Register(s)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error")
	}

	go http.Serve(l, nil)
	<-c

}
