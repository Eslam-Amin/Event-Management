package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
func InitDB(){
	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		fmt.Println(err)
		panic("Couldn't connect to database")
	}
	
	DB.SetMaxIdleConns(10)
	DB.SetMaxIdleConns(5)
}

