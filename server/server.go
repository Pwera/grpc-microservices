package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/pwera/gRPC-notes/todo"
	grpc "google.golang.org/grpc"
)

type server struct{}

func (*server) Unary(ctx context.Context, req *todo.GreetRequest) (*todo.GreetResponse, error) {
	fmt.Println("Unary function called %w", req)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("The server cancelled the request")
			return nil, status.Error(codes.Canceled, "The server cancelled the request")
		}
		time.Sleep(1 * time.Second)
	}
	first := req.GetGreet().GetFirst()
	if first == "" {
		return nil, status.Errorf(
			codes.InvalidArgument, "Received an empty string",
		)
	}
	result := "Hello" + first
	res := &todo.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) ServerStreaming(req *todo.GreetRequest, stream todo.GreetService_ServerStreamingServer) error {
	fmt.Println("ServerStreaming function called %w", req)
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
func (*server) ClientStreaming(stream todo.GreetService_ClientStreamingServer) error {
	fmt.Println("ClientStreaming function called")
	result := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&todo.GreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatal("Error while reading client %w", err)
		}

		first := req.GetGreet().GetFirst()
		result += first + " "
		fmt.Println(result)
	}

}

func (*server) StreamEveryone(stream todo.GreetService_StreamEveryoneServer) error {
	fmt.Println("StreamEveryone function called")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal("Error while reading client stream %w", err)
			return err
		}
		first := req.GetGreet().GetFirst()
		result := "Hello " + first + " !"
		err = stream.Send(&todo.GreetResponse{
			Result: result,
		})
		if err != nil {
			log.Fatal("Error while sending client stream %w", err)
			return err
		}
	}

}

func main() {
	fmt.Println("Server has started")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	tls := true
	opts := []grpc.ServerOption{}
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed loading certificates %v", sslErr)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	srv := grpc.NewServer(opts...)
	todo.RegisterGreetServiceServer(srv, &server{})

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
