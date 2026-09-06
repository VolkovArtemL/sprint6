package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "MORSE-SERVER: ", log.Ldate|log.Ltime)

	srv := server.New(logger)

	logger.Println("сервер запущен на порту 8080")
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
