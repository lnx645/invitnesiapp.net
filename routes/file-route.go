package routes

import (
	"embed"
	"io/fs"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func InitPublicDirectory(app *fiber.App, dir embed.FS) {
	strippedFS, err := fs.Sub(dir, "public")
	if err != nil {
		log.Fatal(err)
	}
	//ini adalah get data yang ada di folder public
	app.Get("assets/*", static.New("", static.Config{
		FS: strippedFS,
	}))
}
