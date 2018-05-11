package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "demo")
	result, _ := ioutil.ReadAll(r.Body)
	w.Write(result)
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	os.Exit(0)
}
