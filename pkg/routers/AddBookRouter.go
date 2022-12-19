package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/pkg/models"
)

func (r Routers) AddBookRouter(c *fiber.Ctx) error {
	var book models.Books

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    book,
	})
}
