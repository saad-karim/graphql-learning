package database

import (
	"fmt"

	"github.com/pkg/errors"
)

type Event struct {
	ID             string `db:"id"`
	Type           string `db:"eventable_type"`
	EventID        string `db:"eventable_id"`
	Status         string `db:"status"`
	Reason         string `db:"reason"`
	Message        string `db:"message"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
	OrganizationID string `db:"organization_id"`
}

type DB interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

type EventDatabase struct {
	DB
}

func (e *EventDatabase) GetAll() ([]Event, error) {
	fmt.Println("!!SK >>> EventDatabase - GetAll()")

	var event []Event

	if err := e.DB.Select(&event, "Select * FROM events WHERE id = 1943 OR id = 4983"); err != nil {
		return nil, errors.Wrap(err, "failed to complete get query")
	}

	return event, nil
}

func (e *EventDatabase) Get(id string) (*Event, error) {
	fmt.Println("!!SK >>> EventDatabase - Get()")

	var event Event

	if err := e.DB.Get(&event, "Select * FROM events WHERE id = $1", id); err != nil {
		return nil, errors.Wrap(err, "failed to complete get query")
	}

	return &event, nil
}
