package entities

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var mydb *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:root@tcp(118.89.50.110:3306)/youhome?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	mydb = db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
