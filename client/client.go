package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	// doServerStreaming(c)

	// doClientStreaming(c)

	doBiDiStreaming(c)
}

func doBiDiStreaming(c todo.GreetServiceClient) {
	fmt.Printf("Start BidiStreaming %v", c)
	//create  a stream by invoking the client
	stream, err := c.StreamEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}
	requests := []*todo.GreetRequest{
		&todo.GreetRequest{
			Greet: &todo.Greeting{
				First:  "First1",
				Second: "Second1",
			},
		},
		&todo.GreetRequest{
			Greet: &todo.Greeting{
				First:  "First2",
				Second: "Second2",
			},
		},
	}

	waitc := make(chan struct{})
	// send a bunch of messages
	go func() {
		for _, req := range requests {
			fmt.Println("Sending data %w", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// receive a bunch of messages
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				close(waitc)
			}
			fmt.Println("Received: %w\n\n", res.GetResult())
		}
	}()
	<-waitc
}

func doClientStreaming(c todo.GreetServiceClient) {
	fmt.Printf("Start ClientStreaming %v", c)

	requests := []*todo.GreetRequest{
		&todo.GreetRequest{
			Greet: &todo.Greeting{
				First:  "First1",
				Second: "Second1",
			},
		},
		&todo.GreetRequest{
			Greet: &todo.Greeting{
				First:  "First2",
				Second: "Second2",
			},
		},
	}

	stream, err := c.ClientStreaming(context.Background())
	if err != nil {
		log.Fatal("error while calling doClientStreaming %w", err)
	}

	for _, req := range requests {
		fmt.Println("Sending data...")
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("error while receiving response: %w", err)
	}
	fmt.Println(res)

}

func doServerStreaming(c todo.GreetServiceClient) {
	fmt.Printf("Start ServerStreaming %v", c)
	req := &todo.GreetRequest{
		Greet: &todo.Greeting{
			First:  "Pioter",
			Second: "Wera",
		},
	}
	resStream, err := c.ServerStreaming(context.Background(), req)
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
	res, err := c.Unary(context.Background(), req)
	if err != nil {
		log.Fatalf("Couldn't connect to server %v", err)
	}
	fmt.Printf("Response from server %v", res)
}
