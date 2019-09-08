package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/pwera/gRPC-notes/src/main/go/todo"
	"google.golang.org/grpc"
)

func main() {
	tls := false
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt"
		creds, sslError := credentials.NewClientTLSFromFile(certFile, "")
		if sslError != nil {
			log.Fatalf("Eror while loading CA trust certificate: %v", sslError)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Couldnt connect to: %v", err)
	}
	defer conn.Close()
	c := todo.NewGreetServiceClient(conn)

	doUnary(c)

	doServerStreaming(c)

	doClientStreaming(c)

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
	requests := prepareData()

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

	requests := prepareData()

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
	req := prepareData()[0]
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	res, err := c.Unary(ctx, req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr)
			switch respErr.Code() {
			case codes.InvalidArgument:
				fmt.Println("ERROR: Empty string sended")
				break
			case codes.DeadlineExceeded:
				fmt.Println("ERROR: Deadline was exceeded")
				break
			default:
				log.Fatalf("Couldn't connect to server %v", err)
				break
			}

		} else {
			log.Fatalf("Erro while calling service %v", err)
		}
		return
	}
	fmt.Printf("Response from server %v", res)
}

func prepareData() []*todo.GreetRequest {
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
	fmt.Println("prepareData: %w", requests)
	return requests
}
