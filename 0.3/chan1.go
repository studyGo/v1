package main

import (
	"fmt"
)

type typeRes struct {
	name string
	t    string
	host string
	port int
	key  string
}

type channel struct {
	c chan string
}

func (t *typeRes) Redis() string {

	return "link redis and send data"
}

func (t *typeRes) Rab() {

}

func main() {
	t := &typeRes{"s", "a", "127.0.0.1", 6379, "rds"}
	m := make(map[string]*channel)
	m[t.name] = &channel{make(chan string, 1)}
	go func() {
		data := t.Redis()
		m[t.name].c <- data
	}()

	fmt.Println(<-m[t.name].c)
}
