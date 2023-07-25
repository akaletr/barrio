package main

import (
	"log"

	"cmd/main/main.go/internal/server"
)

func main() {
	srv := server.New()
	log.Fatal(srv.Start())
}
