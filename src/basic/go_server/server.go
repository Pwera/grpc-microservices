package main

import (
	"context"
	"fmt"
	"github.com/pwera/basicgoserver/basicpb"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) MyServiceUnary(ctx context.Context, req *basicpb.BasicRequest) (*basicpb.BasicResponse, error) {
	fmt.Println("Unary function called %w", req)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("The server cancelled the request")
			return nil, status.Error(codes.Canceled, "The server cancelled the request")
		}
		time.Sleep(1 * time.Second)
	}
	first := req.Value
	if first == "" {
		return nil, status.Errorf(
			codes.InvalidArgument, "Received an empty string",
		)
	}
	res := &basicpb.BasicResponse{
		Id:     0,
		Status: "?",
	}
	return res, nil
}

func (*server) MyServiceServerStreaming(req *basicpb.BasicRequest, stream basicpb.BasicService_MyServiceServerStreamingServer) error {
	fmt.Println("ServerStreaming function called %w", req)
	first := req.Value
	for i := 0; i < 10; i++ {

		result := "Hello" + first + strconv.Itoa(i)
		res := &basicpb.BasicResponse{
			Id:     0,
			Status: result,
		}
		stream.Send(res)
		time.Sleep(100 * time.Millisecond)
	}
	return nil

}
func (*server) MyServiceClientStreaming(stream basicpb.BasicService_MyServiceClientStreamingServer) error {
	fmt.Println("ClientStreaming function called")
	result := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&basicpb.BasicResponse{
				Status: result,
				Id:     0,
			})
		}
		if err != nil {
			log.Fatal("Error while reading client %w", err)
		}

		first := req.Value
		result += first + " "
		fmt.Println(result)
	}

}

func (*server) MyServiceBiDiStreaming(stream basicpb.BasicService_MyServiceBiDiStreamingServer) error {
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
		err = stream.Send(&basicpb.BasicResponse{
			Id:     req.Id,
			Status: req.String(),
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

	opts := []grpc.ServerOption{}

	srv := grpc.NewServer(opts...)
	s := server{}
	basicpb.RegisterBasicServiceServer(srv, &s)

	reflection.Register(srv)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
