package models

import "time"

type Event struct {
	ID int
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string	`binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int 
	CreatedAt time.Time
}

var events = []Event{} 

func New() *Event{

	return &Event{}
}