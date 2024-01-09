package main

import (
	"context"
	"echoStars/greetings"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	greetings.UnimplementedGreeterServiceGrpcServer
	greeterService greetings.GreeterService
}

func (s *server) Hello(ctx context.Context, in *greetings.HelloRequest) (*greetings.HelloResponse, error) {
	log.Printf("Received: %v", in.Message)
	return &greetings.HelloResponse{Message: s.greeterService.Hello(in.Message)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetings.RegisterGreeterServiceGrpcServer(s, &server{
		greeterService: greetings.NewGreeterService(greetings.GreeterImp{}),
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
