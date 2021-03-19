package graph

import "github.com/saad-karim/graphql-learning/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type EventDatabase interface {
	Get(id string) (*database.Event, error)
	GetAll() ([]database.Event, error)
}

type Resolver struct {
	Event EventDatabase
}
