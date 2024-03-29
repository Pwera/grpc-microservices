// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: order.proto

package grpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateOrderRequest) Validate() error {
	if !(this.UserId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must be greater than '0'`, this.UserId))
	}
	if !(this.UserId < 10000) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must be less than '10000'`, this.UserId))
	}
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}
func (this *Item) Validate() error {
	if this.ProductCode == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProductCode", fmt.Errorf(`value '%v' must not be an empty string`, this.ProductCode))
	}
	return nil
}
func (this *CreateOrderResponse) Validate() error {
	return nil
}
