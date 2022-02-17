#!/bin/sh

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/tutorial/tutorial.proto \
    proto/app1/app1.proto
