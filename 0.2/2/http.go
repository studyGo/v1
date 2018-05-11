package main

import (
	"fmt"
	"io"
	"net/http"
)

type ht struct {
}

func (h *ht) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.create(w, r)
}

func (h *ht) create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "this is a demo get")
	} else {
		io.WriteString(w, "this is a demo post")
	}
	fmt.Println(1)
}

func main() {
	ttt := &ht{}
	http.ListenAndServe(":10101", ttt)

}
