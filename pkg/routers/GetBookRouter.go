package routers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/pkg/models"
)

func (r Routers) GetBookRouter(c *fiber.Ctx) error {
	id := c.Params("id")

	var book = []models.Books{
		{
			ID:     1,
			Title:  "The Lord of the Rings",
			Author: "J.R.R. Tolkien",
			Year:   1954,
		},
		{
			ID:     2,
			Title:  "The Hobbit",
			Author: "J.R.R. Tolkien",
			Year:   1954,
		},
	}

	for _, v := range book {
		int_id, err := strconv.Atoi(id)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Invalid ID",
			})

		}

		if v.ID == int_id {
			return c.JSON(fiber.Map{
				"success": true,
				"data":    v,
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Book not found",
	})
}
