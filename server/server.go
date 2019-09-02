package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/pwera/gRPC-notes/todo"
	grpc "google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *todo.GreetRequest) (*todo.GreetResponse, error) {
	fmt.Println("Greet function called %w", req)
	first := req.GetGreet().GetFirst()
	result := "Hello" + first
	res := &todo.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *todo.GreetRequest, stream todo.GreetService_GreetManyTimesServer) error {
	fmt.Println("GreetManyTimes function called %w", req)
	first := req.GetGreet().GetFirst()
	for i := 0; i < 10; i++ {

		result := "Hello" + first + strconv.Itoa(i)
		res := &todo.GreetResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(100 * time.Millisecond)
	}
	return nil

}

func main() {
	fmt.Println("Server has started")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	srv := grpc.NewServer()

	todo.RegisterGreetServiceServer(srv, &server{})

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
