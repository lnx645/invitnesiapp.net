package models

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	ID        uint         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
	DeletedAt sql.NullTime `gorm:"index" json:"deleted_at,omitempty"`
}
