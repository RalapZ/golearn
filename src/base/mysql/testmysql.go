package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var mysqlstr = "xconf:123456@tcp(10.10.8.150:3306)/xconf"

type User struct {
	Id   int
	Name string
	Sex  string
}

var conn *sql.DB

func InitDB(sqlstr string) {
	var err error
	conn, err = sql.Open("mysql", sqlstr)
	if err != nil {
		fmt.Println(err)
	}
	err = conn.Ping()
	if err != nil {
		fmt.Println(err)
	}
	conn.SetMaxIdleConns(10)
}

func main() {
	InitDB(mysqlstr)
	sqlstring := "INSERT INTO user(id,name,sex) VALUES(2,'ralap1','male1')"
	conn.Exec(sqlstring)

	r, err1 := conn.Query(" select * from user;")
	//conn.QueryRow()
	if err1 != nil {
		fmt.Println(err1)
	}
	a := User{}
	for r.Next() {
		r.Scan(&a.Id, &a.Name, &a.Sex)
		fmt.Println(a)
	}
}
