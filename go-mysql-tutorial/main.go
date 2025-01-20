package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users VALUES('hoge')")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Successfully inserted user tables")
}
