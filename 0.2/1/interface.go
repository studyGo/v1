package main

import "fmt"

type demo interface {
    get() string
    post() string
}

type s struct {
    name string
}

func (sl s) get() string {
    return  sl.name
}

func (sl s) post() string {

    return sl.name
}


func main() {
    dl := s{"demo"}
    fmt.Println(dl.get())
    tl := s{"post"}
    fmt.Println(tl.post())
}

