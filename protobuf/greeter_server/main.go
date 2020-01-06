/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../protobuf/ --go_out=plugins=grpc:../protobuf/ ../protobuf/helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/nalcheg/copy-paste-go/protobuf/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// SayEhlo implements helloworld.GreeterServer
func (s *server) SayEhlo(ctx context.Context, in *pb.EhloRequest) (*pb.EhloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.EhloReply{Message: "Ehlo " + in.GetName()}, nil
}

func (s *server) SaySixSixSix(ctx context.Context, in *pb.SaySixSixSixRequest) (*pb.SaySixSixSixReply, error) {
	return &pb.SaySixSixSixReply{SixSixSixInt: 666, SixSixSixString: "666"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
