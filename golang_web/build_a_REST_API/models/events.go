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
	// thay k can thiet nen k nhan result
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	// thay k can thiet nen cmt
	// id, err := result.LastInsertId()
	// e.ID = id

	// result.LastInsertId()

	return err
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, errors.New("Can't reading data!")
	}

	return &event, nil
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

func UpdateEvents(id int64, event *Event) error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.ID)
	if err != nil {
		return errors.New("Failed to exec query")
	}
	// van chua hieu vi sao truyen thieu user id ma van cahy dc
	return nil
}

func (event Event) DeleteEvent() error {
	query := `
	DELETE FROM events 
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("Couldn't prepare stmt")
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)

	if err != nil {
		return errors.New("Couldn't execute DELETE statement")
	}

	return nil
}
