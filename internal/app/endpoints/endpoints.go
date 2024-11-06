package endpoints

import (
	"myApp/internal/app/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Endpoints struct {
	s *service.Service
}

func New() *Endpoints {
	return &Endpoints{}
}

func (e *Endpoints) Login(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func (e *Endpoints) Index(c *fiber.Ctx) error {
	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"status": "ok",
	})
}
