package main

import (
	"flag"
	"fmt"
	"gopkg.in/redis.v2"
	"time"
)

type typeRes struct {
	name string
	t    string
	host string
	port int
	key  string
}

type data struct {
	name string
	data string
}

func main() {
	var path = flag.String("p", "./chan.conf", "config path")
	// read conf
	count, conf := parseConf(args)
	for i := 0; i < count; i++ {
		switch conf[i].(*typeRes).t {
		case redis:
			go createRedis()
		case rab:
			go createRab()
		default:
		}
	}
	for {
		readData()
	}
}

func readData() {
	// link rab
	// parse data
	/*
			   if data == nil {
			       time.Sleep(time.Second)
			   }
		       <-



	*/
}

func (t *typeRes) createRedis() {
	// link readis
	client := redis.NewTCPClient(&redis.Options{
		Addr:     t.host + ":" + t.port,
		Password: "",
		DB:       0,
	})

	if client != nil {
		fmt.Println(t.name + " is not connect")
	}
	var s data

	for {
		s = <-getData

		go func() {
			action := make(chan bool, 1)
			timeoutC := make(chan bool, 1)
			go func() {
				client.LPush(t.key, t.data)
				action <- true
			}()
			go func() {
				timeout(1)
				timeoutC <- true
			}()

			select {
			case <-action:
				returnData <- true

			case <-timeoutC:
				returnData <- false
			}
		}()

		if <-returnData {
			fmt.Println("exec errro")
		}
	}
}

func (t *typeRes) createRab() {

}

func timeout(t int) {
	time.Sleep(time.Second * t)
}

func parseConf(path string) (int []*typeRes) {

}
