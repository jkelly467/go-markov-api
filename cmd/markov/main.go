package main

import (
	"context"
	"log"

	"github.com/mb-14/gomarkov"
	api "go-markov-api"
	"go-markov-api/gomarkovlib"
	"go-markov-api/server"
)

func main() {
	ctx := context.Background()

	markovs := map[string]api.Markov{
		"two":   &gomarkovlib.Markov{*gomarkov.NewChain(2)},
		"three": &gomarkovlib.Markov{*gomarkov.NewChain(3)},
	}

	srv := &server.Server{
		Markovs: markovs,
		Port:    "3000",
	}

	log.Println("Server starting on port: 3000")
	if err := srv.Start(ctx); err != nil {
		log.Fatalf("server.Start: %v", err)
	}
}
