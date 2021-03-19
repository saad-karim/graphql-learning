package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/saad-karim/graphql-learning/graph/generated"
	"github.com/saad-karim/graphql-learning/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	fmt.Println("!!SK >>> queryResolver - Events()")

	var mevents []*model.Event

	events, err := r.Resolver.Event.GetAll()
	if err != nil {
		return nil, err
	}

	for _, event := range events {
		mevents = append(mevents, &model.Event{
			ID:     event.ID,
			Type:   event.Type,
			Status: event.Status,
		})
	}
	// dummyLink := model.Event{
	// 	ID:     "1",
	// 	Status: "Status1",
	// 	Type:   "Disk",
	// }
	// events = append(events, &dummyLink)
	// dummyLink2 := model.Event{
	// 	ID:     "2",
	// 	Status: "Status2",
	// 	Type:   "Disk",
	// }
	// events = append(events, &dummyLink2)
	return mevents, nil
}

func (r *queryResolver) Event(ctx context.Context, id string) (*model.Event, error) {
	fmt.Println("!!SK >>> queryResolver - Event()")

	event, err := r.Resolver.Event.Get(id)
	if err != nil {
		return nil, err
	}

	return &model.Event{
		ID:     event.ID,
		Type:   event.Type,
		Status: event.Status,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
