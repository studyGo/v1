package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

type DB struct {
	res *sql.DB
}

func (d *DB) Conn() {
	u := "root"
	p := "123456"
	db, _ := sql.Open("mysql", u+":"+p+"@/network_sell")
	d.res = db
}

var d DB

func main() {
	d.Conn()
	s()
	fmt.Println(d.res.Query("SHOW TABLES;"))
}
func s() {

	fmt.Println(d.res.Query("SHOW TABLES;"))
}
