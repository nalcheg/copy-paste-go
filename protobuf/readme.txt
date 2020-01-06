go run greeter_server/main.go
go run greeter_client/main.go

protoc -I ../protobuf/ --go_out=plugins=grpc:../protobuf/ ../protobuf/helloworld/helloworld.proto
