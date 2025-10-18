package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"` // lets gin know these fields must be in the request body, so bindwithJSON cannot auto set them to null.  sends an error instead if empty
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e *Event) Save() {
	// later: add to db
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events
}
