package main

import(
    "fmt"
)

type Lock struct {
    read bool
    write bool
}

func New() *Lock{
    return &Lock{
        read : false,
        write : false,
    }
}

func (l *Lock) Read () {
    l.read = true
}

func (l *Lock) UnRead () {
    l.read = false
}

func (l *Lock) Write () {
    l.write = true
}

func (l *Lock) UnWrite () {
    l.write = false
}

func main () {

}
