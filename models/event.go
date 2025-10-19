package models

import (
	"time"

	"github.com/MichaelVenturi/go-practice-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"` // lets gin know these fields must be in the request body, so bindwithJSON cannot auto set them to null.  sends an error instead if empty
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	// later: add to db
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES(?, ?, ?, ?, ?) 
	`
	// te ?s are to protect against sql injection attacks.  they will be filled in via the exec method
	sqlstmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer sqlstmt.Close()
	res, err := sqlstmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	e.ID = id
	return err
}

func (e *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	sqlstmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer sqlstmt.Close()
	_, err = sqlstmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	return err
}

func (e *Event) Delete() error {
	query := `
	DELETE FROM events WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}

func (e *Event) Register(userId int64) error {
	query := `
	INSERT INTO registrations(event_id, user_id) VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `
	DELETE FROM registrations WHERE event_id = ? AND user_id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM events
	`
	rows, err := db.DB.Query(query) // exec for queries changing database, query for queries just fetching
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `
	SELECT * FROM events WHERE id = ?
	`
	row := db.DB.QueryRow(query, id) // gives back one row, once again providing args to fill in the ? mark.

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
