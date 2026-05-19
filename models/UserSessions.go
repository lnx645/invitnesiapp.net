package models

import (
	"time"

	"gorm.io/gorm"
)

type UserSession struct {
	gorm.Model
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	Token     string    `gorm:"unique;not null;size:225" json:"token"`
	ExpiredAt time.Time `gorm:"index" json:"expired_at"`
	UserAgent string    `json:"user_agent"`
	IPAddress string    `json:"ip_address" gorm:"type:varchar(50);index"`
	User      *User     `gorm:"foreignKey:UserID"  json:"user"`
}
