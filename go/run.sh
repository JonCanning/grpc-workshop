#!/bin/bash
protoc --go_out=. --go-grpc_out=. -I=../proto hello.proto
go run .