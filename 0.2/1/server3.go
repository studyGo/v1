package main

import (
	"io"
	"math/rand"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	data := "this is a demo"
	n := rand.Intn(4096) + 1024
	buffer := make([]int, n)
	for i := 0; i < n; i++ {
		buffer[i] = rand.Intn(1024)
	}
	c := 0
	for i := 0; i < n; i++ {
		if buffer[i] > 512 {
			c += 1
		}
	}

	w.Header().Add("Can-You-Help-Me", "yes")
	io.WriteString(w, data)
}

func main() {
	var server http.Server
	http.HandleFunc("/", handle)
	server.Addr = "127.0.0.1:8080"
	server.ListenAndServe()
}
