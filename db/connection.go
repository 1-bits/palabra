package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/palabradb")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
