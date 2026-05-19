package models

import "gorm.io/gorm"

type User struct {
	Name        string        `json:"name"`
	Email       string        `gorm:"size:20;uniqueIndex;not null" json:"email"`
	Password    string        `json:"password"`
	PhoneNumber string        `gorm:"uniqueIndex;not null;size:16" json:"phone_number"`
	IsActive    bool          `json:"is_active"`
	IsProPlan   bool          `json:"is_pro_plan"`
	Sessions    []UserSession `gorm:"foreignKey:UserID"`
	gorm.Model
}
