package routes

import (
	"invitnesia/api/controllers"
	"invitnesia/api/lib"

	"github.com/gofiber/fiber/v3"
)

func AuthRoute(app *fiber.App) {

	authController := controllers.AuthController{
		DB: lib.DB,
	}

	api := app.Group("/api")
	api.Get("/login", authController.Login).Name("auth.login")
}
