package routes

import (
	"invitnesia/api/app/controllers"
	"invitnesia/api/app/lib"
	"invitnesia/api/app/middleware"

	"github.com/gofiber/fiber/v3"
)

func AuthRoute(app *fiber.App) {

	authController := controllers.AuthController{
		DB: lib.DB,
	}

	api := app.Group("/api")
	api.Get("/login", middleware.ProtectedRoute, authController.Login).Name("auth.login")

	v1 := api.Group("/v1")

	v1.Get("/login", middleware.ProtectedRoute, authController.Login).Name("auth.login.v1")

}
