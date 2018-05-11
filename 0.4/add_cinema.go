package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	var u = flag.String("u", "root", "db user name")
	var p = flag.String("p", "netsale", "db pass word")
	var code = flag.String("code", "10000000", "code")
	var ip = flag.String("ip", "8.8.8.8", "ip")

	flag.Parse()

	db, err := sql.Open("mysql", *u+":"+*p+"@/network_sell_api")
	if err != nil {
		fmt.Println("can't connect mysql")
		os.Exit(0)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("can't connect mysql")
		os.Exit(0)
	}

	if !addCinema(db, *code, *ip) {
		fmt.Println("cinema is exists")
		os.Exit(0)
	}
	addCronInfo(db, *code)
}

func addCinema(db *sql.DB, code string, ip string) bool {

	sql := "SELECT `code` FROM netsale_cinema WHERE code = " + code

	res, err := db.Query(sql)

	if err != nil {
		return false
	}
	res.Next()
	var c int
	res.Scan(&c)

	if c != 0 {
		return false
	}

	sql = "INSERT INTO `netsale_cinema` (`code`, `version`, `username`, `password`, `soap_uri`,`public_key`,`sign`, `name`, `address`) VALUES('" + code + "', '1.0', 'netsale', '8e0e6fce8fd33ba3774bfb479a21403b', 'https://" + ip + ":9000/index.net.php?wsdl', '','','','')"
	db.Exec(sql)
	return true
}

func addCronInfo(db *sql.DB, code string) {
	sql := "INSERT INTO netsale_cron_info(`cinema`, `action`, `time`) VALUES(" + code + ", 'cinema', 86400)"
	db.Exec(sql)
	sql = "INSERT INTO netsale_cron_info(`cinema`, `action`, `time`) VALUES(" + code + ", 'seat', 86400)"
	db.Exec(sql)
	sql = "INSERT INTO netsale_cron_info(`cinema`, `action`, `time`) VALUES(" + code + ", 'film', 10800)"
	db.Exec(sql)
	sql = "INSERT INTO netsale_cron_info(`cinema`, `action`, `time`) VALUES(" + code + ", 'session', 3600)"
	db.Exec(sql)
	sql = "INSERT INTO netsale_cron_info(`cinema`, `action`, `time`) VALUES(" + code + ", 'sessionSeat', 240)"
	db.Exec(sql)
	sql = "INSERT INTO netsale_cron_info(`cinema`, `action`, `time`) VALUES(" + code + ", 'screen', 86400)"
	db.Exec(sql)
}

func usage() {
	var usageTemplate = `             
    Usage:
    add_agent [args]

    args
        -u  dbuser 
        -p  dbpassword
        -code  cinema code
        -ip  cinema ip
        -help
    `
	fmt.Println(usageTemplate)
}

//end
