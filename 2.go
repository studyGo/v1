package main

import (
	"fmt"
)

type M struct {
	name string
}

func (m *M) Read(p string) {
	fmt.Println(p)
}

func (m *M) Write(p string) {
	fmt.Println(p)
}

type R interface {
	Read(p string)
}

type W interface {
	Write(p string)
}

func main() {
	m := &M{name: "kalasa"}
	var s R
	s = m
	s.Read("ksd")

	var t W
	t = m
	t.Write("hh")
}
