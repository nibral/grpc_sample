package main

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Service struct {
	id string
}

func (service *Service) SayHello(ctx context.Context, message *HelloRequest) (*HelloReply, error) {
	return &HelloReply{
		Msg: "Hello from " + service.id,
	}, nil
}

func main() {
	port, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	helloService := &Service{
		id: uuid.New().String(),
	}
	RegisterHelloServer(server, helloService)
	err = server.Serve(port)
	if err != nil {
		log.Fatalln(err)
	}
}
