package main

import (
	"log"

	"github.com/shawlyelshalawy/logsync/internal/server"
)

func main() {

	srv := server.NewHttpServer(":3000")
	log.Fatal(srv.ListenAndServe())
}
