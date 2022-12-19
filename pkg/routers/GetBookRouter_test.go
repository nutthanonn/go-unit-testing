package routers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	utils "github.com/gofiber/fiber/v2/utils"
)

func Test_GetBook(t *testing.T) {
	app := fiber.New()
	router := NewRouter()

	app.Get("/books/:id", router.GetBookRouter)

	t.Run("FAIL_PARAMS", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/books/aasndf", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		utils.AssertEqual(t, nil, err, "app.Test(req)")
		utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode, "Status code")
	})

	t.Run("CANT_FIND_BOOK", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/books/10", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		utils.AssertEqual(t, nil, err, "app.Test(req)")
		utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
	})
}
