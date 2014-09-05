package main

import "fmt"
import "time"

// 着玩意可以写计划任务 哈哈哈哈
func main() {

    ticker := time.NewTicker(time.Second)

    go func() {
       for t := range ticker.C {
           fmt.Println(t)
       }
    }()

    timer := time.NewTimer(10 * time.Second)
    <-timer.C
    ticker.Stop()
}

