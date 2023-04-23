package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Username  string         `json:"username" gorm:"type:varchar(30);uniqueIndex:uidx_username"`
	Password  string         `json:"-" gorm:"type:varchar"` // hashing
	Email     string         `json:"email" gorm:"type:varchar(50);"`
	Role      string         `json:"role" gorm:"type:varchar;"`
	Status    string         `json:"status" gorm:"type:varchar;"`
}
