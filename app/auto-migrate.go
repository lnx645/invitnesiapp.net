package app

import (
	"invitnesia/api/lib"
	"invitnesia/api/models"
)

func RunAutoMigrate() {
	lib.DB.AutoMigrate(&models.User{}, &models.UserSession{})
}
