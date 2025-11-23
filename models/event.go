package models

import (
	"time"

	"example.com/event-booking/db"
)

type Event struct {
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string	`binding:"required"`
	EventDate time.Time `binding:"required"`
	UserID int 
	CreatedAt time.Time
}


func (event *Event) Save() error {

    if event.CreatedAt.IsZero() {
        event.CreatedAt = time.Now().UTC()
    }

    query := `
    INSERT INTO events
    (name, description, location, event_date, user_id, created_at)
    VALUES (?, ?, ?, ?, ?, ?);
    `

    stmt, err := db.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    res, err := stmt.Exec(
        event.Name,
        event.Description,
        event.Location,
        event.EventDate,
        event.UserID,
        event.CreatedAt,
    )
    if err != nil {
        return err
    }

    id, err := res.LastInsertId()
    if err != nil {
        return err
    }

    event.ID = id
    return nil
}


func GetAllEvents() ([]Event, error){
	query := `SELECT * from events`
	rows, err := db.DB.Query(query)

	if err != nil{
		return nil, err
	}
	
	var events []Event
	
	defer rows.Close()
	
	for rows.Next(){
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.EventDate, &event.UserID, &event.CreatedAt)
		
		if err != nil{
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func New() *Event{

	return &Event{}
}