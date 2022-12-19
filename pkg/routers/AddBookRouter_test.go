package routers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	utils "github.com/gofiber/fiber/v2/utils"
)

func Test_AddBook(t *testing.T) {
	app := fiber.New()
	router := NewRouter()
	app.Post("/books", router.AddBookRouter)

	t.Run("FAIL_POST_JSON", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/books", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		utils.AssertEqual(t, nil, err, "app.Test(req)")
		utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode, "Status code")
	})
}
