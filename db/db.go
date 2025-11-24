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
	createDBTables()
}


func createDBTables(){
	createEventsTable()
	createUsersTable()
}

func createUsersTable(){
	query := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`

	_, err := DB.Exec(query)
	if err != nil{
		fmt.Println("err in creating users table ",err)
		panic(err)
	}
}

func createEventsTable(){
query := `CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		event_date DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		created_at DATETIME NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id) 
		);`
		
		
		_, err := DB.Exec(query) 
		if err != nil {
		fmt.Println("err in creating events table ",err)
		panic(err)
	}
}