package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var mysqlstr = "xconf:123456@tcp(10.10.8.150:3306)/xconf"

func main() {
	db, err := sql.Open("mysql", mysqlstr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
	db.SetMaxIdleConns(10)
	r, _ := db.Query(" show tables;")
	fmt.Println(r.Scan())

}
