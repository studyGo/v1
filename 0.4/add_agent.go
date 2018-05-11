package main

//
//  用来增加 第三方待售信息
//
import (
	"crypto/md5"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func main() {
	var u = flag.String("u", "", "username")
	var p = flag.String("p", "", "password")
	var name = flag.String("n", "", "app_name")
	var cname = flag.String("cn", "", "company name")
	var tel = flag.String("t", "", "telphone")
	var con = flag.String("c", "", "contacts")
	var help = flag.String("help", "", "help")
	var dbUser = flag.String("dbu", "root", "db_user")
	var dbPass = flag.String("dbp", "netsale", "db_pass")
	var dbHost = flag.String("dbh", "127.0.0.1", "db_Host")
	flag.Parse()
	//fmt.Println(*u, p, name, cname, tel, con, *help)
	if (*u == "" || *p == "" || *name == "" || *cname == "" || *tel == "" || *con == "") && *help == "" {
		fmt.Println("params is not empty")
		os.Exit(0)
	}
	db, err := sql.Open("mysql", *dbUser+":"+*dbPass+"@tcp("+*dbHost+":3306)/network_sell_api")
	if err != nil {
		fmt.Println("can't connect mysql")
		os.Exit(0)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("can't connect mysql")
		os.Exit(0)
	}

	pass := getMd5(*p)
	token := getToken()
	times := fmt.Sprintf("%d", time.Now().Unix())

	sql := "INSERT INTO `netsale_app`(`username`, `password`, `access_token`, `create_time`, `app_name`, `company_name`, `telphone`, `contacts`) VALUES ('" + *u + "','" + pass + "','" + token + "','" + times + "','" + *name + "','" + *cname + "','" + *tel + "', '" + *con + "')"
	db.Exec(sql)

}

func usage() {
	var usageTemplate = `             
    Usage:
    add_agent [args]

    args
        -u  username   
        -p  epassword
        -n  app name
        -cn company_name
        -t  telphone
        -c  contacts
        -dbu  dbUser
        -dbp  dbPass
        -dbh  dbHost
        -help
    `
	fmt.Println(usageTemplate)
}

func getToken() string {
	//	h := md5.New()
	str := time.Now().String()
	data := []byte(str)
	d := md5.Sum(data)
	return fmt.Sprintf("%x", d)
}

func getMd5(str string) string {
	data := []byte(str)
	d := md5.Sum(data)
	return fmt.Sprintf("%x", d)
}
