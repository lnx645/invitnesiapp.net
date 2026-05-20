package controllers

import (
	"fmt"
	"invitnesia/api/app/lib"
	"invitnesia/api/app/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func CreateAuthController() {

}

func (c *AuthController) Login(ctx fiber.Ctx) error {
	var product []models.User
	data := lib.RedisClient.Get(ctx.Context(), "auth:name").Val()
	fmt.Println(data)
	err := c.DB.Model(&models.User{}).Preload("Sessions").Find(&product).Error
	if err != nil {
		return err
	}
	return ctx.JSON(product)
}
