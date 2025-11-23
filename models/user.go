package models

type User struct {
	ID int 
	Name string `binding:"required"`
	Email string `binding:"required"`
	Password string `binding:"required"`
}


func NewUser()*User{
	return &User{}
}