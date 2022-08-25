package main

import (
	"context"
	"log"
	"shapes/server"
)

func main() {
	if err := server.NewHTTPServer().Run(context.Background()); err != nil {
		log.Fatalf("failed starting server, err : \n%+v", err)
	}
}
