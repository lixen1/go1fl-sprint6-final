package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)

	s := server.NewServer(logger)

	logger.Println("Серевер запущен")
	err := s.HTTPServer.ListenAndServe()
	if err != nil {
		logger.Fatal("ошибка во время работы сервера: ", err)
	}
}
