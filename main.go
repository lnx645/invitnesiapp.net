package main

import (
	"embed"
	"invitnesia/api/app"
	"invitnesia/api/config"
	"invitnesia/api/lib"
	"invitnesia/api/middleware"
	"invitnesia/api/routes"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
)

//go:embed public/*
var publicDIR embed.FS

func main() {
	quit := make(chan os.Signal, 1)
	//database init
	err := lib.InitDBConnection(&config.Get().Database)
	//redis init
	if err != nil {
		log.Fatal("Gagal Koneksi ke server Database mysql: " + err.Error())

	}

	if err := lib.InitRedisClient(&config.Get().Redis); err != nil {
		log.Fatal("Gagal Koneksi ke server Redis: " + err.Error())
	}
	app.RunAutoMigrate()
	fiberConfig := config.FiberConfig()
	appServer := fiber.New(*fiberConfig)

	middleware.FiberMiddleware(appServer)
	routes.MainRoute(appServer)
	routes.AuthRoute(appServer)
	routes.InitPublicDirectory(appServer, publicDIR)
	appServer.Use(func(ctx fiber.Ctx) error {
		ctx.SendStatus(fiber.StatusNotFound)
		if ctx.IsJSON() || ctx.Accepts("json", "html") == "json" {
			return ctx.JSON(fiber.Map{
				"status":  "404",
				"message": "Waduh, nyasar ya? ",
			})
		}
		return ctx.Render("404", fiber.Map{})
	})
	go func() {
		log.Println("Server Invitnesia berjalan di port :8080")
		if err := appServer.Listen(":8080"); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Mematikan server secara perlahan (Graceful Shutdown)...")

	if err := appServer.Shutdown(); err != nil {
		log.Fatalf("Error during shutdown: %v", err)
	}
	log.Println("Server Invitnesia berhasil dimatikan dengan aman.")
}
