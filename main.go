package main

import (
	"log"
	"text-classifier/api"
)

func main() {
	server := api.NewServer()
	log.Fatal(server.Start(":8080"))
}
