package grpc

import (
	"google.golang.org/grpc"
	"log"
)

type ClientFunc func(TodoServiceClient, interface{})

func DoClient (f ClientFunc, option interface{}){
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := NewTodoServiceClient(cc)
	f(c, option)
}
