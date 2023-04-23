package models

import (
	"gorm.io/gorm"
	"time"
)

// Article ...
type Article struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
	UserID    uint           `json:"-"`
	User      *User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
}
