package routes

import "github.com/gofiber/fiber/v3"

func AuthRoute(app *fiber.App) {
	api := app.Group("/api")
	//api login
	api.Get("/login", func(ctx fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	}).Name("auth.login")
}
