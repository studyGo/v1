package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	if os.Getppid() != 1 {
		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)
		//将其他命令传入生成出的进程
		cmd.Stdin = os.Stdin //给新进程设置文件描述符，可以重定向到文件中
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start() //开始执行新进程，不等待新进程退出
		return
	} else {
		time.Sleep(time.Second * 100)
	}
}
