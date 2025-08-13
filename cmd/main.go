package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stderr, "", log.LstdFlags)
	srv := server.New(logger)
	if err := srv.HTTP.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
