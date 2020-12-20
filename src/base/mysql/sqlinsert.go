package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User1 struct {
	Id   int
	Name string
	Sex  string
}

var conn1 *sql.DB
var sqldriver = "mysql"
var sqlstr = "xconf:123456@tcp(10.10.8.150:3306)/xconf"

func InittestDB() {
	var err error
	conn1, err = sql.Open(sqldriver, sqlstr)
	if err != nil {
		fmt.Println(err)
	}
	err = conn1.Ping()
	if err != nil {
		fmt.Println(err)
	}
}

func Execsql(sqlstr string) {
	res, err := conn1.Exec(sqlstr)
	if err != nil {
		fmt.Println(err)
	}
	i, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	k, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(k, i)
}

func Presqlinsert(sqlinsert string, a User1) {
	exepre, err := conn1.Prepare(sqlinsert)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	exepre.Exec(a.Id, a.Name, a.Sex)
}

func main() {
	InittestDB()
	//sqlinsert:="insert into user(id,name,sex)values(4,'test','female');"
	sqlinsert2 := "insert into user(id,name,sex)values(?,?,?);"
	a := User1{10, "zhu", "male"}
	fmt.Println(a)
	//Execsql(sqlinsert)
	Presqlinsert(sqlinsert2, a)
}
