package todo

import (
	"github.com/kkmmttdd/bff-graphql-to-grpc-golang/client/grpc"
	graph "github.com/kkmmttdd/bff-graphql-to-grpc-golang/server/graph/model"
)

func GraphToGrpc(graph_t *graph.Todo) *grpc.Todo{
	return &grpc.Todo{Text:graph_t.Text}
}

func GrpcToGraph(grpc_t *grpc.Todo) *graph.Todo{
	return &graph.Todo{Text:grpc_t.Text}
}
