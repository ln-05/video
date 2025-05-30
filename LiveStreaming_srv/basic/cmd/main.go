package main

import (
	"LiveStreaming_srv/basic/inits"
	"LiveStreaming_srv/hander"
	__ "LiveStreaming_srv/proto"
	"flag"
	"fmt"
	"google.golang.org/grpc"

	"log"
	"net"
)

func main() {
	inits.Init()
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8300))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterUserServer(s, &hander.UserServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
