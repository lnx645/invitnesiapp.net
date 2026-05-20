package routes

import (
	"fmt"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
)

type JsonData struct {
	RequestId string `json:"requestId"`
	Path      string `json:"path"`
	BaseUrl   string `json:"baseUrl"`
	ApiUrl    string `json:"apiUrl"`
}

func MainRoute(app *fiber.App) {
	app.Get("app/*", func(ctx fiber.Ctx) error {
		if strings.Contains(ctx.Path(), ".") {
			return ctx.Next()
		}
		data := JsonData{
			RequestId: ctx.RequestID(),
			Path:      ctx.Path(),
			BaseUrl:   ctx.BaseURL(),
			ApiUrl:    ctx.Path("/api"),
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			return ctx.Status(500).SendString("Error parsing JSON")
		}
		return ctx.Render("app", fiber.Map{
			"InjectedJson": string(jsonData),
			"ApiUrl":       fmt.Sprintf("%s/%s", ctx.BaseURL(), "api"),
			"AssetUrl":     fmt.Sprintf("%s/%s", ctx.BaseURL(), "assets"),
		})
	})

}
