package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	httpServer := server.CreateNewServer(logger)

	err := httpServer.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
