#!/bin/sh

protoc --proto_path=. \
       --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       common.proto bridge.proto gw.proto
