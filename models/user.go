package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
	Username  string         `json:"username" gorm:"type:varchar(30);uniqueIndex:uidx_username"`
	Password  string         `json:"-" gorm:"type:varchar"` // hashing
	Email     string         `json:"email" gorm:"type:varchar(50);"`
	Role      string         `json:"role" gorm:"type:varchar;"`
	Status    string         `json:"status" gorm:"type:varchar;"`
}
