package main

import "fmt"
import "runtime"
import "math/rand"
import "time"
import "sync"

var total_ticket int32 = 100;
var mutex = &sync.Mutex{}

func sell_ticket(i int) {
    for total_ticket > 0{
        mutex.Lock()
        if total_ticket > 0 {
            //time.Sleep(time.Duration(rand.Intn(5))* time.Millisecond)
            total_ticket--
            fmt.Println(i ,"sell ticket: total", total_ticket)
        }

        mutex.Unlock()
    }


}

func main() {
    runtime.GOMAXPROCS(4)

    rand.Seed(time.Now().Unix())
    for i := 0; i< 5; i++ {
        go sell_ticket(i)
    }

    var input string
    fmt.Scanln(&input)
}
