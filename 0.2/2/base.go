package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var str = "this is a demo"
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(str)))

}

func Auth(username string, password string) string {
	auth = username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
