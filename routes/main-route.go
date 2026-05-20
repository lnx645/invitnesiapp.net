package routes

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

func MainRoute(app *fiber.App) {
	app.Get("app/*", func(ctx fiber.Ctx) error {
		if strings.Contains(ctx.Path(), ".") {
			return ctx.Next()
		}
		return ctx.Render("app", fiber.Map{})
	})

}
