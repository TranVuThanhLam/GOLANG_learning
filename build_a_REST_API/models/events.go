package models

import (
	"errors"
	"time"

	"example.com/web/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"` // khong cach va phai co dau ""
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)
	`
	// da tung ghi dau , thanh dau . :))

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	// da tung thieu user ID
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id

	// result.LastInsertId()

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var events []Event

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, errors.New("Can't reading data!")
		}

		events = append(events, event)
	}
	return events, nil
}
