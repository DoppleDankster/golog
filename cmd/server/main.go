package main

import (
	"log"

	"github.com/DoppleDankster/golog/internal/server"
)

func main() {

	srv := server.NewHTTPServer(":8000")
	log.Fatal(srv.ListenAndServe())

}
