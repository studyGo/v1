package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var data string
	if req.Method == "POST" {
		fmt.Println(req)
		req.ParseMultipartForm(32 << 20)
		if req.MultipartForm != nil {
			values := req.MultipartForm.Value["data"]
			data = values[0]
		}
		log(data)
	}
	io.WriteString(w, data)
}

func log(data string) {

	file := "test.log"
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("openfile error")
	}
	f.WriteString("[" + data + "] \n")
}

func main() {
	var server http.Server
	http.HandleFunc("/log", handle)
	server.Addr = "127.0.0.1:8080"
	server.ListenAndServe()
}
