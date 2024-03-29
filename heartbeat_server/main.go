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

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/WiscEdgeCentralController/heartbeat"
	"google.golang.org/grpc"
)

const (
	port = ":50050"
)

var allClientInfo map[string]clientInfo

// server is used to implement helloworld.GreeterServer.
type server struct{}

type clientInfo struct{
	clientId string
	meaasge string
}



// SayHello implements helloworld.GreeterServer
func (s *server) ReceiveAndReply(ctx context.Context, in *pb.HeartbeatRequest) (*pb.HeartbeatReply, error) {
	clientId := in.ClientId
	info, ok := allClientInfo[clientId]

	if ok {
		//update info
		info.meaasge = in.Name
		log.Printf("Received: %v", in.Name)
	} else {
		//creat new client instance
		allClientInfo[clientId] = clientInfo{clientId, in.Name}
	}
	return &pb.HeartbeatReply{Message: "Hello " + in.Name}, nil
}

func main() {
	allClientInfo= make(map[string]clientInfo)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHeartbeatPBServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
