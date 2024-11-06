package main

import (
	"log"
	"myApp/internal/pkg/app"
	"os"

	"github.com/joho/godotenv"
)

var (
	port string
)

func init() {
	if err := godotenv.Load("port.env"); err != nil {
		log.Print("No .env file found")
	}
	var exists bool
	port, exists = os.LookupEnv("PORT")

	if !exists {
		log.Fatal("Ошибка обралотки .env", port)
	}
}

func main() {
	app, err := app.New()

	if err != nil {
		log.Fatal("Server can't starting")
	}

	app.Run(port)
}
