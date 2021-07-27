# About

Example of using ReactJS + Go + gRPC


Install required Go plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

## How to generate proto

From the project's root folder run:

```bash
  protoc --go_out=server/ --go-grpc_out=server/ protos/hello.proto
```
  


