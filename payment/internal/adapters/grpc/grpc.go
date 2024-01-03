package grpc

import (
	"context"
	"fmt"
	"github.com/pwera/grpc-micros-payment/internal/application/core/domain"
	"github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"

	//"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	logrus.Info("Payments::Create")
	if false {
		badReq := &errdetails.BadRequest{}
		orderStatus := status.New(codes.InvalidArgument, "something went wrong")
		statusWithDetails, _ := orderStatus.WithDetails(badReq)
		return nil, statusWithDetails.Err()
	}

	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &CreatePaymentResponse{PaymentId: result.ID}, nil
}
