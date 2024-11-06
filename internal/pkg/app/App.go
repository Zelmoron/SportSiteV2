package app

import (
	"SportSite/internal/app/endpoints"
	"SportSite/internal/app/services"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	endpoins *endpoints.Endpoints
	services *services.Services
	app      *fiber.App
}

func MiddleWare(c *fiber.Ctx) error {
	value := c.Get("User-status")

	if value == "admin" {
		c.Locals("status", value)
		return c.Next()
	}
	return c.SendString("Доступ запрещен")
}

func New() *App {
	a := &App{}
	a.app = fiber.New()

	a.services = services.New()
	a.endpoins = endpoints.New(a.services)

	a.app.Get("/", MiddleWare, a.endpoins.Index)
	a.app.Post("/login", a.endpoins.Login)

	return a

}

func (a *App) Run() {
	a.app.Listen(":8080")
}
