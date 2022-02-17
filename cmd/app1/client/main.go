// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/panlw/using-grpc/proto/app1"
)

const (
	defaultName = "Neo"
)

var (
	addr = flag.String("addr", "localhost:17001", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewApp1ProtoRpcClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	hi(ctx, c)
	hello(ctx, c)
}

func hi(ctx context.Context, c pb.App1ProtoRpcClient) {
	r, err := c.Greet(ctx, &pb.GreetReq{Greeting: *name, Mode: pb.EchoMode_hi})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r.GetHi().Who)
}

func hello(ctx context.Context, c pb.App1ProtoRpcClient) {
	r, err := c.Greet(ctx, &pb.GreetReq{Greeting: *name, Mode: pb.EchoMode_hello})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r.GetHello().How)
}
