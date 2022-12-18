package main

import (
	"go-microservices-training/services/router"
	"log"
)

func main() {
	log.Println("Starting this application")
	router.StartSv("8080")
}