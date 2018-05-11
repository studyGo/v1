package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//var s map[string][]string = make(map[string][]string)
	var AppPath string

	app := filepath.Join(AppPath, "Conf", "Config")

	fmt.Println(app)
	fmt.Println(os.Getwd())
}
