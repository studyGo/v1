package main

import (
    "fmt"
    "flag"
)

func main () {
    var ip = flag.Int("flagname", 100085, "help message for flagname")
    flag.Parse()
    fmt.Println(*ip)

}
