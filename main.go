package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/PongDev/SW-Arch-File-Storage-Microservice/app"
	"github.com/PongDev/SW-Arch-File-Storage-Microservice/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.SetupEnvironment()

	s, err := app.NewServer()
	if err != nil {
		panic(err)
	}

	go func() {
		s.Start(config.Env.PORT)
	}()

	shutdownSignal := make(chan os.Signal, 1)

	signal.Notify(shutdownSignal, os.Interrupt)

	<-shutdownSignal

	s.GracefulStop()
	if err := s.Cleanup(); err != nil {
		log.Fatal(err)
	}
}
