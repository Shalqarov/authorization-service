package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Shalqarov/authorization-service/auth"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "server port")

type helloServer struct {
	pb.UnimplementedAuthorizationServer
}

func (server *helloServer) Hello(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Reply{Str: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthorizationServer(s, &helloServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
