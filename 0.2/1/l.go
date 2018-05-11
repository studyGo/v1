package main

import (
    "fmt"
    "sync"
    "time"
)

var m *sync.Mutex

func main () {
    m = new(sync.Mutex)
    go lock(3)
    time.Sleep(time.Second)
    lock(1)
}
func lock (i int) {
    m.Lock()
    time.Sleep(time.Second * 10)
    fmt.Println("demo", i)
    m.Unlock()
}
