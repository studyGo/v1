package main

import "fmt"
import "math/rand"
import "time"

func main() {
    //eb06a63d-cfe9-4211-ac35-d67c80d44339
    var msg string
    s := "0123456789abcdef"
    rand.Seed(time.Now().Unix())
    msg = ""
    for i := 0; i < 8; i++ {
        d := rand.Intn(15)
        msg = msg + s[d:d+1]

    }
    msg = msg + "-"
    for i := 0; i < 4; i++ {
        d := rand.Intn(15)
        msg = msg + s[d:d+1]

    }
    msg = msg + "-"
    for i := 0; i < 4; i++ {
        d := rand.Intn(15)
        msg = msg + s[d:d+1]

    }
    msg = msg + "-"
    for i := 0; i < 4; i++ {
        d := rand.Intn(15)
        msg = msg + s[d:d+1]

    }
    msg = msg + "-"
    for i := 0; i < 12; i++ {
        d := rand.Intn(15)
        msg = msg + s[d:d+1]

    }
    fmt.Println(msg)
}
