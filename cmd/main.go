package main

import (
	"go-microservices-training/services/router"
	"log"
)

func main() {
	log.Println("Starting this application")
	host:=":8080"
	router.StartH2CServer(host)
}