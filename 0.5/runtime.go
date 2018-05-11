package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	fmt.Println(file, apppath)

}
