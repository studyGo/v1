package main

import (
	"fmt"
	"log"
	"net/http"
)

type M struct {
	logger *log.Logger
	action http.Handler
}

func New() *M {
	m := &M{}
	return m
}

func HelloServe(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(rw)
}

func (m *M) Run() {
	http.HandleFunc("/hello", HelloServe)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func main() {
	New().Run()
}
