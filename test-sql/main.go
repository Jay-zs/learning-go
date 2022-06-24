package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Goyal@921@tcp(127.0.0.1:3306)/test")
	fmt.Println(db.Ping())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success")
	data, err := db.Query("CREATE table employee {}")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(data.Columns())
}
