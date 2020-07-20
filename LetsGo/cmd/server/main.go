package main

import (
	"log"

	"github.com/nydan/proglog/LetsGo/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8000")
	log.Fatal(srv.ListenAndServe())
}
