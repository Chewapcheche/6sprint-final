package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "server:", log.LstdFlags)
	srv := server.NewServer(logger)
	logger.Println("Server start via port 8080")
	if err := srv.HTTP.ListenAndServe(); err != nil {
		logger.Fatal("server start error:", err)
	}
}
