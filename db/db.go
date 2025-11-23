package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB
func InitDB(){
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		fmt.Println("err in db connection ",err)
		panic("Couldn't connect to database")
	}
	
	DB.SetMaxOpenConns(10)
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
		user_id INTEGER NOT NULL
		);`
		
		
		_, err := DB.Exec(createEventsTable) 
		if err != nil {
		fmt.Println("err in creating events table ",err)
		panic(err)
	}
}