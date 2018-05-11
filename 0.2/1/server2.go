package main

import (
	"fmt"
	"io"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)

	data := ""
	w.Header().Add("Can-You-Help-Me", "yes")
	io.WriteString(w, data)
}
func m(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)

	data := "m"
	w.Header().Add("Can-You-Help-Me", "yes")
	io.WriteString(w, data)
	s(w, r)
}
func s(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)

	data := "s"
	w.Header().Add("Can-You-Help-Me", "yes")
	io.WriteString(w, data)
}
func w(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)

	data := "w"
	w.Header().Add("Can-You-Help-Me", "yes")
	io.WriteString(w, data)
}
func t(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)

	data := "t"
	w.Header().Add("Can-You-Help-Me", "yes")
	io.WriteString(w, data)
}

func main() {
	var server http.Server
	http.HandleFunc("/m", m)
	http.HandleFunc("/s", s)
	http.HandleFunc("/w", w)
	http.HandleFunc("/t", t)
	server.Addr = "127.0.0.1:8080"
	server.ListenAndServe()
}
