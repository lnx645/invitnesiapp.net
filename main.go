package main

import (
	"embed"
	"invitnesia/api/app"
	"invitnesia/api/config"
	"invitnesia/api/lib"
	"invitnesia/api/middleware"
	"invitnesia/api/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

//go:embed public/*
var publicDIR embed.FS

func main() {
	//database init
	err := lib.InitDBConnection(&config.Get().Database)
	//redis init
	lib.InitRedisClient(&config.Get().Redis)
	if err != nil {
		panic(err.Error())
	}
	app.RunAutoMigrate()
	fiberConfig := config.FiberConfig()
	app := fiber.New(*fiberConfig)

	middleware.FiberMiddleware(app)
	routes.MainRoute(app)
	routes.AuthRoute(app)
	routes.InitPublicDirectory(app, publicDIR)
	app.Use(func(ctx fiber.Ctx) error {
		ctx.SendStatus(fiber.StatusNotFound)
		return ctx.Render("404", fiber.Map{})
	})
	log.Fatal(app.Listen(":8080"))
}
