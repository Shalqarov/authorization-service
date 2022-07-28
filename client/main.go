package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/Shalqarov/authorization-service/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "World", "Name to say hello :)")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthorizationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.Hello(ctx, &pb.Request{Name: *name})
	if err != nil {
		log.Fatalf("could not say hello: %v", err)
	}
	log.Println(response.GetStr())
}
