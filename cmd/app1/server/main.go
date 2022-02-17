package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/panlw/using-grpc/proto/app1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	port = flag.Int("port", 17001, "The server port")
)

// server is used to implement App1ProtoRpcServer
type server struct {
	pb.App1ProtoRpcServer
}

// Greet implements App1ProtoRpcServer#Greet
func (s *server) Greet(ctx context.Context, req *pb.GreetReq) (*pb.GreetRes, error) {
	mode := req.GetMode()
	log.Printf("[APP1] Mode: %v", mode)
	if mode == pb.EchoMode_hi {
		return &pb.GreetRes{Greeting: &pb.GreetRes_Hi_{Hi: &pb.GreetRes_Hi{Who: "Neo"}}}, nil
	}
	if mode == pb.EchoMode_hello {
		return &pb.GreetRes{Greeting: &pb.GreetRes_Hello_{Hello: &pb.GreetRes_Hello{How: "quickly"}}}, nil
	}
	return nil, errors.New("NG_MODE")
}

var empty = &emptypb.Empty{}

// Write implements App1ProtoRpcServer#Write
func (s *server) Write(ctx context.Context, req *pb.WriteReq) (*emptypb.Empty, error) {
	fmt.Printf("[APP1] Content: %s\n", req.Content)
	return empty, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterApp1ProtoRpcServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
