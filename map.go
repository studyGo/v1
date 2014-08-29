package main

import "fmt"

func main() {
    m := make(map[string]bool)

    m["demo"] = false
    m["fuck"] = true

    fmt.Println(m)
    fmt.Println(m["demo"])
    for key, val := range m {
        fmt.Println(key, val)
    
    }
}
