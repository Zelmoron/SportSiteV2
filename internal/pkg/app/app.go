package app

import (
	"log"
	"myApp/internal/app/endpoints"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	endpoints *endpoints.Endpoints
	fiber     *fiber.App
}

func New() (*App, error) {
	app := &App{}

	app.fiber = fiber.New()
	app.endpoints = endpoints.New()

	app.fiber.Post("/login", app.endpoints.Login)
	app.fiber.Get("/", app.endpoints.Index)

	return app, nil
}

func (app *App) Run(port string) {
	log.Fatal(app.fiber.Listen(port))
}
