package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kkmmttdd/bff-graphql-to-grpc-golang/server/graph/generated"
	"github.com/kkmmttdd/bff-graphql-to-grpc-golang/server/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var links []*model.Todo
	dummyTodo := model.Todo{
		ID: "hogehoge",
		Text: "Dummy",
		Done:false,
		User: &model.User{
			ID:"hogehgoe",
			Name: "namename",
		},
	}
	todos := append(links, &dummyTodo)
	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
