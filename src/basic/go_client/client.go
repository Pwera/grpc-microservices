package main

import (
	"context"
	"flag"
	"fmt"
	basicpb2 "github.com/pwera/basicgoclient/basicpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func main() {
	var server_host string
	var server_port string
	flag.StringVar(&server_host, "server_host", "127.0.0.1", "")
	flag.StringVar(&server_port, "server_port", "50051", "")
	flag.Parse()
	opts := grpc.WithInsecure()

	conn_url := server_host + ":" + server_port
	fmt.Println("Go client trying connect to %w", conn_url)
	conn, err := grpc.Dial(conn_url, opts)
	if err != nil {
		log.Fatalf("Couldnt connect to: %v", err)
	}
	defer conn.Close()
	c := basicpb2.NewBasicServiceClient(conn)

	doUnary(c)

	doServerStreaming(c)

	doClientStreaming(c)

	doBiDiStreaming(c)
}

func doBiDiStreaming(c basicpb2.BasicServiceClient) {
	fmt.Printf("Start BidiStreaming %v", c)
	//create  a stream by invoking the client
	stream, err := c.MyServiceBiDiStreaming(context.Background())
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
			fmt.Println("Received: %w\n\n", res)
		}
	}()
	<-waitc
}

func doClientStreaming(c basicpb2.BasicServiceClient) {
	fmt.Printf("Start ClientStreaming %v", c)

	requests := prepareData()

	stream, err := c.MyServiceClientStreaming(context.Background())
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

func doServerStreaming(c basicpb2.BasicServiceClient) {
	fmt.Printf("Start ServerStreaming %v", c)
	req := prepareData()[0]

	resStream, err := c.MyServiceServerStreaming(context.Background(), req)
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
		log.Printf("response from server: %v", msg)
	}

}

func doUnary(c basicpb2.BasicServiceClient) {
	fmt.Printf("Start Unary %v", c)
	req := prepareData()[0]
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	res, err := c.MyServiceUnary(ctx, req)
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

func prepareData() []*basicpb2.BasicRequest {
	requests := []*basicpb2.BasicRequest{
		{
			Id:    0,
			Value: "0",
		},
		{
			Id:    1,
			Value: "1",
		},
	}
	return requests
}
