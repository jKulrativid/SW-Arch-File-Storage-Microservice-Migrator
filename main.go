package main

import (
	"github.com/PongDev/Go-gRPC-Storage/app"
	"github.com/PongDev/Go-gRPC-Storage/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.SetupEnvironment()

	s := app.NewServer()
	s.Start(config.Env.PORT)
}
