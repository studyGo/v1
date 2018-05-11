package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	var reply int

	if err != nil {
		fmt.Println(err)
	}

	err = cli.Call("S.Multiply", "ss", &reply)
	fmt.Println(reply)
	var get string
	err = cli.Call("S.Get", "sd", &get)
	fmt.Println(get)

}
