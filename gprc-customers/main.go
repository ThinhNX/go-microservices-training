package main

import (
	"fmt"
	"go-microservices-training/gprc-customers/handlers"
	"go-microservices-training/gprc-customers/repositories"
	"go-microservices-training/helper"
	"go-microservices-training/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := helper.AutoBindConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	listen, err := net.Listen("tcp", ":2222")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	customerRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewCustomerHandler(customerRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterCustomerServiceServer(s, h)
	fmt.Println("Listen on port 2222")
	s.Serve(listen)
}
