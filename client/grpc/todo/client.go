package todo

import (
	"context"
	"github.com/kkmmttdd/bff-graphql-to-grpc-golang/client/grpc"
	"log"
)

type ListOption struct {
	*grpc.TodoListRequest
	*grpc.TodoListResponse
}

func List(c grpc.TodoServiceClient, option interface{}) {
	opt := option.(ListOption)
	r, err := c.TodoList(context.Background(), opt.TodoListRequest)
	println(r.Todos[0].Text)
	//println((*r).Todos[0].Text)
	*opt.TodoListResponse = *r
	println(opt.Todos[0].Text)
	if err != nil {
		log.Fatal(err)
	}
}



