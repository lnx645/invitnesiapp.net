package middleware

import (
	"invitnesia/api/app/lib"

	"github.com/gofiber/fiber/v3"
)

func ProtectedRoute(ctx fiber.Ctx) error {
	//ini nanti ambil semua data dari frontend
	lib.RedisClient.Set(ctx.Context(), "auth:100283", "John Doe", 0)

	return ctx.Next()
}
