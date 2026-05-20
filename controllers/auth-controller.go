package controllers

import (
	"invitnesia/api/models"

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
	err := c.DB.Model(&models.User{}).Preload("Sessions").Find(&product).Error
	if err != nil {
		return err
	}
	return ctx.JSON(product)
}
