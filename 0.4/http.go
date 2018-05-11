/*
   use http timeouthandler

*/
package main

import (
	"io"
	"net/http"
	"time"
)

func handle(rw http.ResponseWriter, r *http.Request) {

}

type HTT struct{}

func (htt *HTT) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 2)
	io.WriteString(rw, "this is a demo")

}

func main() {

	s := &HTT{}
	//http.HandleFunc("/he", handle)
	//http.Handle("/he", s)
	http.Handle("/he", http.TimeoutHandler(s, time.Second, "Timeout"))
	http.ListenAndServe(":2222", nil)

}
