package main

import (
	"fmt"
	"os"
)

func main() {

	file := "test.log"
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	//func (f *File) WriteString(s string) (ret int, err error)
	f.WriteString("this is a demo")
	//func (f *File) Write(b []byte) (n int, err error)
	_, err = f.Write([]byte("this is a demo"))
	if err != nil {
		fmt.Println(err)
	}

	fStat, _ := os.Stat(file)
	fmt.Println(fStat.Name())
	fmt.Println(fStat.Size())
	fmt.Println(fStat.Mode())
	fmt.Println(fStat.ModTime())
	fmt.Println(fStat.IsDir())
	fmt.Println(fStat.Sys())

	f, err = os.Open(file)
	buffer := make([]byte, 100)
	_, err = f.Read(buffer)
	fmt.Println(buffer)

}
