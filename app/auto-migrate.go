package app

import (
	"invitnesia/api/app/lib"
	"invitnesia/api/app/models"
)

func RunAutoMigrate() {
	lib.DB.AutoMigrate(&models.User{}, &models.UserSession{})
}
