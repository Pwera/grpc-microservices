#!/usr/bin/env bash

## protoc-gen-govalidators need to be on PATH

order=../order/internal/adapters/grpc
payment=../payment/internal/adapters/grpc

## Order
protoc \
--proto_path=$GOPATH/pkg/mod/github.com/mwitkow/go-proto-validators@v0.3.2 \
--proto_path=. \
--go_out=$order \
--go_opt=paths=source_relative \
--go-grpc_out=$order \
--go-grpc_opt=paths=source_relative \
--govalidators_out=$order \
--govalidators_opt=paths=source_relative  \
order.proto


protoc \
--proto_path=$GOPATH/pkg/mod/github.com/mwitkow/go-proto-validators@v0.3.2 \
--proto_path=. \
--go_out=$order \
--go_opt=paths=source_relative \
--go-grpc_out=$order \
--go-grpc_opt=paths=source_relative \
--govalidators_out=$order \
--govalidators_opt=paths=source_relative  \
payment.proto


## Payments
protoc \
--proto_path=$GOPATH/pkg/mod/github.com/mwitkow/go-proto-validators@v0.3.2 \
--proto_path=. \
--go_out=$payment \
--go_opt=paths=source_relative \
--go-grpc_out=$payment \
--go-grpc_opt=paths=source_relative \
--govalidators_out=$payment \
--govalidators_opt=paths=source_relative  \
payment.proto