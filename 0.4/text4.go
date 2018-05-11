package main

/**
  使用hashtable 减少每次任务循环次数
*/

import (
	"bytes"
	"container/list"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/redis.v2"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const max = 2 << 4
const filmTypeTime, filmPriceTime, releaseTime, updateTime, cinemaStatus = 3600 * 3, 3600 * 3, 60, 3 * 60, 5 * 60
const phpBin = "/usr/local/m1905/php/bin/php"
const phpFile = "/data/html/netsale_api/index.cron.php"
const tmpFile = "/data/logs/cron/gocron"
const timeoutSecond = 1200

var mysqlCronDB DB
var redisCron RC
var cronT = &CronProvider{list: list.New()}
var item = make([]*Items, max)

type Items struct {
	l *list.List
}

type CronProvider struct {
	list *list.List
}

type CronList struct {
	time   int64
	action string
	cinema int
	delay  int64
}

func getIndex(val int) int {
	return val % max

}

type DB struct {
	res *sql.DB
}

func (d *DB) Conn(u string, p string) {
	db, err := sql.Open("mysql", u+":"+p+"@/network_sell_api")
	if err != nil {
		fmt.Println("can't connect mysql")
		os.Exit(0)
	}
	d.res = db
}

type RC struct {
	res *redis.Client
}

func (d *RC) Conn() {
	client := redis.NewTCPClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	d.res = client
}

func execute() {
	res, err := mysqlCronDB.res.Query("SELECT i.* FROM netsale_cron_info i LEFT JOIN netsale_cinema c ON i.cinema = c.code where c.is_show = 1")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var tDelay int64 = 20
	unixTime := time.Now().Unix() + tDelay

	index := getIndex(int(time.Now().Unix() + tDelay))
	item[index] = &Items{list.New()}

	//影片类型
	item[index].l.PushFront(&CronList{unixTime, "type", 0, filmTypeTime})
	//第三方价格
	item[index].l.PushBack(&CronList{unixTime, "price", 0, filmPriceTime})
	//自动解锁
	item[index].l.PushBack(&CronList{unixTime, "release", 0, releaseTime})
	item[index].l.PushBack(&CronList{unixTime, "update", 0, updateTime})
	//清除redis老数据
	item[index].l.PushBack(&CronList{unixTime, "clean", 0, filmPriceTime})
	//特殊状态 自动出售
	item[index].l.PushBack(&CronList{unixTime, "sell", 0, updateTime})
	//获取已出售的二维码
	item[index].l.PushBack(&CronList{unixTime, "qrget", 0, updateTime})

	for res.Next() {
		var id int
		var cinema int
		var action string
		var delay int64

		res.Scan(&id, &cinema, &action, &delay)
		item[index].l.PushBack(&CronList{unixTime, action, cinema, delay})
	}

	res, err = mysqlCronDB.res.Query("SELECT code FROM netsale_cinema where is_show = 1")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for res.Next() {
		var code int
		res.Scan(&code)

		// 影院网络状态
		item[index].l.PushBack(&CronList{unixTime, "status", code, cinemaStatus})
	}

}

func do(actionLen chan int, channel chan string) {

	for {
		time.Sleep(time.Millisecond * 800)
		unixTime := time.Now().Unix()
		index := getIndex(int(unixTime))

		if item[index] == nil {
			continue
		}

		// 循环hashtable 链表
		for s := item[index].l.Front(); s != nil; s = s.Next() {

			if s.Value.(*CronList).time <= time.Now().Unix() {
				cronlist := s.Value.(*CronList)
				actionLen <- 1
				item[index].l.Remove(s)

				go func() {
					runType := run(cronlist, channel)
					if runType {
						index := getIndex(int(unixTime + cronlist.delay))
						if item[index] == nil {
							item[index] = &Items{list.New()}
						}
						item[index].l.PushBack(&CronList{unixTime + cronlist.delay, cronlist.action, cronlist.cinema, cronlist.delay})
					} else {
						index := getIndex(int(unixTime + 3))
						if item[index] == nil {
							item[index] = &Items{list.New()}
						}
						item[index].l.PushBack(&CronList{time.Now().Unix() + 3, cronlist.action, cronlist.cinema, cronlist.delay})
					}
					<-actionLen
				}()
			}
		}
	}
}

func main() {
	var u = flag.String("u", "root", "db user name")
	var p = flag.String("p", "netsale", "db pass word")
	var pt = flag.Int("pt", 30, "Parallel task Num")
	var help = flag.String("help", "", "help")
	flag.Parse()

	if *help != "" {
		usage()
		os.Exit(0)
	}

	mysqlCronDB.Conn(*u, *p)
	redisCron.Conn()
	channel := make(chan string, *pt)
	actionLen := make(chan int, *pt)

	execute()

	// 记录log
	go func() {
		f, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
		if err != nil {
			os.Exit(0)
		}
		for {
			data := <-channel
			f.WriteString(time.Now().String() + "   " + data + "\n")
		}
	}()

	do(actionLen, channel)
}

func usage() {
	var usageTemplate = `             
    Usage:
    cron [args]

    args
        -u  db username   
        -p  db password
        -pt Parallel task Num
        -help
    `
	fmt.Println(usageTemplate)
}

func log() {
	redisCron.res.Incr("report:error")
}

func run(cron *CronList, c chan string) bool {

	var e bool

	t := time.Now().UnixNano()

	switch cron.action {
	case "cinema":
		e = cinema(cron.cinema)
		break
	case "screen":
		e = screen(cron.cinema)
		break
	case "seat":
		e = seat(cron.cinema)
		break
	case "film":
		e = film(cron.cinema)
		break
	case "session":
		e = session(cron.cinema)
		break
	case "sessionSeat":
		e = sessionSeat(cron.cinema)
		break
	case "type":
		e = filmType()
		break
	case "price":
		e = filmPrice()
		break
	case "release":
		e = releaseOrder()
		break
	case "update":
		e = updateOrder()
		break
	case "clean":
		e = clean()
		break
	case "sell":
		e = sell()
		break
	case "status":
		e = status(cron.cinema)
		break
	case "qrget":
		e = qrcode()
	default:
		break
	}

	o := strconv.FormatInt((time.Now().UnixNano()-t)/1000000, 10)
	o = o + "ms"

	if e {
		c <- cron.action + " " + strconv.Itoa(cron.cinema) + " SUCCESS " + o
	} else {
		c <- cron.action + " " + strconv.Itoa(cron.cinema) + " FAILES " + o
		go log()
	}
	return e
}

func action(a string, cinema int) bool {
	var out bytes.Buffer
	var timeoutChan = make(chan bool, 1)
	var execChan = make(chan bool, 1)
	var pidChan = make(chan int, 1)
	var returnBool bool

	go func() {
		e := exec.Command(phpBin, phpFile, a, strconv.Itoa(cinema))
		e.Stdout = &out
		e.Start()
		pid := e.Process.Pid
		pidChan <- pid
		e.Wait()
		if out.String() != "Success" {
			execChan <- false
		}

		execChan <- true
		close(execChan)
	}()

	go func() {
		time.Sleep(time.Second * timeoutSecond)
		timeoutChan <- true
	}()

	select {
	case <-timeoutChan:
		pid := <-pidChan
		p, _ := os.FindProcess(pid)
		p.Kill()
	case returnBool = <-execChan:
		return returnBool
	}

	close(timeoutChan)
	close(pidChan)
	return false

}

func cinema(cinema int) bool {
	return action("Cron/Cinema/get", cinema)
}

func screen(cinema int) bool {
	stat := action("Cron/screen/get", cinema)
	if stat {
		return screenType(cinema)
	}
	return false
}

func screenType(cinema int) bool {
	return action("Cron/screen/type", cinema)
}

func seat(cinema int) bool {
	return action("Cron/seat/get", cinema)
}

func film(cinema int) bool {
	return action("Cron/film/get", cinema)
}

func session(cinema int) bool {
	return action("Cron/session/get", cinema)
}

func sessionSeat(cinema int) bool {
	return action("Cron/sessionSeat/get", cinema)
}

func filmType() bool {
	return action("Cron/Film/type", 0)
}

func filmPrice() bool {
	return action("Cron/Film/price", 0)
}

func releaseOrder() bool {
	return action("Cron/UpdateOrder/release", 0)
}

func updateOrder() bool {
	return action("Cron/UpdateOrder/update", 0)
}

func clean() bool {
	return action("Cron/CleanOldData", 0)
}

func sell() bool {
	return action("Cron/UpdateOrder/sell", 0)
}

func status(cinema int) bool {
	return action("Cron/Cinema/status", cinema)
}

func qrcode() bool {
	return action("Cron/QrCode/get", 0)
}
