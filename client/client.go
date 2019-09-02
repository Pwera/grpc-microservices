package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/pwera/gRPC-notes/todo"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldnt connect to: %v", err)
	}
	defer conn.Close()
	c := todo.NewGreetServiceClient(conn)

	// doUnary(c)

	doServerStreaming(c)
}
func doServerStreaming(c todo.GreetServiceClient) {
	fmt.Printf("Start ServerStreaming %v", c)
	req := &todo.GreetRequest{
		Greet: &todo.Greeting{
			First:  "Pioter",
			Second: "Wera",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {

	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream %w", err)
		}
		log.Printf("response from server: %v", msg.Result)
	}

}

func doUnary(c todo.GreetServiceClient) {
	fmt.Printf("Start Unary %v", c)
	req := &todo.GreetRequest{
		Greet: &todo.Greeting{
			First:  "Pioter",
			Second: "Wera",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Couldn't connect to server %v", err)
	}
	fmt.Printf("Response from server %v", res)
}
