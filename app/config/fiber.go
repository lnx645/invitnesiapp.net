package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v3"
)

type AppCtx = fiber.Ctx
type structValidator struct {
	validate *validator.Validate
}

// validate
func (c *structValidator) Validate(out any) error {
	return c.validate.Struct(out)
}

func FiberConfig() *fiber.Config {
	engine := html.New("./views", ".html")
	return &fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
		Views:           engine,
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
		StrictRouting:   true,
		ServerHeader:    "Invitnesia Server",
		AppName:         "Invitnesia Api",
	}
}
