#!/usr/bin/env bash

protoc -I ../proto Basic.proto --go_out=plugins=grpc:basicpb
