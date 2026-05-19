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
	err := lib.InitDBConnection(&config.Get().Database)
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

	log.Fatal(app.Listen(":9000"))
}
