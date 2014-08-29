package main

import "fmt"

type query struct {
    sql chan string
    result chan string
}

func execute(q query) {

    go func() {
        q.result <- "get " + <-q.sql
    }()
}

func main() {
    q := query{make(chan string, 1), make(chan string, 1)}
    execute(q)
    q.sql <- "select * from user"

    fmt.Printf("%s\n", <-q.result)
}
