package app

import (
	"invitnesia/api/lib"
	"invitnesia/api/models"
)

func RunAutoMigrate() {
	lib.DB.AutoMigrate(&models.UserSession{})
	lib.DB.AutoMigrate(&models.Users{})
}
