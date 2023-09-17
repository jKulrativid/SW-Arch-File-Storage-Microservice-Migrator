#!/bin/bash

if [ ! -d "filestorage/grpc" ]; then
  mkdir -p filestorage/grpc
fi
protoc file_storage.proto --go_out=./filestorage/grpc --go_opt=paths=source_relative
protoc file_storage.proto --go-grpc_out=./filestorage/grpc --go-grpc_opt=paths=source_relative
