package main

import (
	"fmt"
	"go-microservices-training/grpc-flights/handlers"
	"go-microservices-training/grpc-flights/repositories"
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
	listen, err := net.Listen("tcp", ":2223")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	flightRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewFlightHandler(flightRepository)
	if err != nil {
		panic(err)
	}
	reflection.Register(s)
	pb.RegisterFlightsServer(s, h)
	fmt.Println("Listen on port 2223")
	s.Serve(listen)
}
