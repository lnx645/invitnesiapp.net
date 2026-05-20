package routes

import "github.com/gofiber/fiber/v3"

func MainRoute(app *fiber.App) {
	app.Get("", func(ctx fiber.Ctx) error {
		return ctx.Render("home", fiber.Map{})
	})

}
