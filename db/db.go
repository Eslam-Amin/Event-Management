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
	createTables()
}


func createTables(){
	createEventsTable := `CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		)`
		
		
		_, err :=DB.Exec(createEventsTable) 
		if err != nil {
		fmt.Println(err)
		panic(err)
	}
}