package models

import (
	"errors"

	"example.com/event-booking/db"
	"example.com/event-booking/utils"
)

type User struct {
	ID int64
	Name string `binding:"required"`
	Email string `binding:"required"`
	Password string `binding:"required"`
}

type LoginCredentials struct {
	Email string `binding:"required"`
	Password string `binding:"required"`
}


func (user *User) Save() error {
	query := `
	INSERT INTO users
	(name, email, password)
	VALUES (?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil{
		return err
	}
	result, err := stmt.Exec(user.Name, user.Email, hashedPassword)
	if err != nil{
		return err
	}
	userId, err := result.LastInsertId()

	user.ID = userId
	return err
}

func (user *User)ValidateCredentials(inputPassword string) error{
	validPassword := utils.ComparePasswords(inputPassword, user.Password)
	
	if !validPassword {
		return errors.New("invalid credentials")
	}

	return nil
}

func GetAllUsers()([]User, error){
	query := `
	Select id, name, email from users;
	`
	users := []User{}
	rows, err := db.DB.Query(query)
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		var user User 
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil{
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByEmail(email string) (*User, error){
	query := `
	select * from users where email = ?;
	`
	var user User
	row := db.DB.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	
return &user, nil
}

func NewUser()*User{
	return &User{}
}