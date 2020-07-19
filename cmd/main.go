package main

import (
	"context"
	"log"

	"github.com/phev8/go_service_B/pkg/server_b"
)

func main() {

	// <---
	ctx := context.Background()

	if err := server_b.RunServer(
		ctx,
		"4302",
	); err != nil {
		log.Fatal(err)
	}

}
