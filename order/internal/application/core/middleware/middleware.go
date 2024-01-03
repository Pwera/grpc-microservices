package middleware

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/sony/gobreaker"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"reflect"
	"time"
)

func createCircuitBreaker(cbName string) *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        cbName,
		MaxRequests: 3,
		Timeout:     4,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio >= 0.6
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("Circuit Breaker: %s, changed from %v, to %v", name, from, to)
		},
	})
}

func CircuitBreakerClientInterceptor(cbName string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		_, cbErr := createCircuitBreaker(cbName).Execute(func() (interface{}, error) {
			err := invoker(ctx, method, req, reply, cc, opts...)
			if err != nil {
				return nil, err
			}

			return nil, nil

		})
		return cbErr
	}
}

func RetryInterceptor() grpc.UnaryClientInterceptor {
	return grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
	)
}

func HandleError(err error, msg string) error {
	if err == nil {
		return nil
	}
	st := status.Convert(err)
	var details []proto.Message
	for _, detail := range st.Details() {
		switch t := detail.(type) {
		case *errdetails.BadRequest:
			details = append(details, t)
		default:
			fmt.Printf("Found an unknown type: %v\n", reflect.TypeOf(t))
			err = status.Newf(codes.Unknown, "Unknown error %s", msg).Err()
		}
	}
	if len(details) > 0 {
		statusWithDetails, _ := status.New(codes.InvalidArgument, msg).
			WithDetails(details...)
		err = statusWithDetails.Err()
	}

	if len(st.Details()) == 0 {
		st, _ := status.FromError(err)
		statusWithDetails, _ := status.New(codes.Unknown, msg).
			WithDetails(&errdetails.ErrorInfo{
				Reason: st.Message(),
			})
		err = statusWithDetails.Err()
	}
	return err
}
