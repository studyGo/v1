package main

import (
	"bytes"
	"container/list"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type CronTimeoutList struct {
	time int64
	pid  int
}
type CronProvider struct {
	list *list.List
}

var cronT = &CronProvider{list: list.New()}

func main() {

	go func() {
		for {
			for s := cronT.list.Front(); s != nil; s = s.Next() {
				crontimeoutlist := s.Value.(*CronTimeoutList)

				if crontimeoutlist.time < time.Now().Unix() {
					p, _ := os.FindProcess(crontimeoutlist.pid)
					p.Kill()
				}
			}
			time.Sleep(time.Second * 2)
		}
	}()

	e := exec.Command("php", "1.php")
	var out bytes.Buffer
	e.Stdout = &out
	e.Start()
	pid := e.Process.Pid
	cronT.list.PushFront(&CronTimeoutList{time.Now().Unix() + 3, pid})
	e.Wait()
	fmt.Println(out.String())
}
