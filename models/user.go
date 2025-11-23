package models

import "example.com/event-booking/db"

type User struct {
	ID int64
	Name string `binding:"required"`
	Email string `binding:"required"`
	Password string `binding:"required"`
}


func (user *User) Save() error {
	query := `
	INSERT TO users
	(name, email, password)
	VALUES (?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil{
		return err
	}
	userId, err := result.LastInsertId()

	user.ID = userId
	return err
}

func NewUser()*User{
	return &User{}
}