package endpoints

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Index(c *fiber.Ctx) string
}

type LoginRequest struct {
	Password string `json:"password"`
}

type Endpoints struct {
	s Service
}

func (e *Endpoints) Index(c *fiber.Ctx) error {
	value := e.s.Index(c)

	return c.SendString("Это главная страница " + value)
}

func (e *Endpoints) Login(c *fiber.Ctx) error {
	var lr LoginRequest
	if err := c.BodyParser(&lr); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Не удалось получить данные")
	}

	if lr.Password == "123" {
		return c.SendString("Пароль правильный")

	}
	return c.SendString("Пароль не правильный")
}

func New(s Service) *Endpoints {
	return &Endpoints{
		s: s,
	}
}
