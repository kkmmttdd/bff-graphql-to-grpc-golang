package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kkmmttdd/bff-graphql-to-grpc-golang/adapter/grpc/todo"
	grpc_client "github.com/kkmmttdd/bff-graphql-to-grpc-golang/client/grpc"
	client "github.com/kkmmttdd/bff-graphql-to-grpc-golang/client/grpc/todo"
	"github.com/kkmmttdd/bff-graphql-to-grpc-golang/server/graph/generated"
	"github.com/kkmmttdd/bff-graphql-to-grpc-golang/server/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TodoList(ctx context.Context) ([]*model.Todo, error) {
	grpcReq := grpc_client.TodoListRequest{}
	grpcRes := grpc_client.TodoListResponse{}
	grpc_client.DoClient(client.List, client.ListOption{
		TodoListRequest:  &grpcReq,
		TodoListResponse: &grpcRes,
	})
	fmt.Println("%d", len(grpcRes.Todos))
	var res []*model.Todo
	for _, gRes := range grpcRes.Todos {
		res = append(res, todo.GrpcToGraph(gRes))
	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
