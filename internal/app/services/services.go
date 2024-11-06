package services

import "github.com/gofiber/fiber/v2"

type Services struct {
}

func (s *Services) Index(c *fiber.Ctx) string {
	return c.Locals("status").(string)
}

func New() *Services {
	return &Services{}
}
