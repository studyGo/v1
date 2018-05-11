package main

import "fmt"
import "container/list"

func main() {
	l := list.New()
	e := l.PushFront("demo")
	e = l.InsertBefore("tt", e)
	e = l.InsertBefore("ss", e)
	l.InsertAfter("ll", e)

	e = l.Front()
	for {
		fmt.Println(e.Value)
		e = e.Next()
		if e == nil {
			return
		}
	}
}
