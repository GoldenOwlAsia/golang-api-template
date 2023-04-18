package gorms

import (
	"time"
)

type User struct {
	Id             uint      `json:"-" gorm:"primaryKey;autoIncrement:true"`
	Username       string    `json:"username" gorm:"type:varchar(30);uniqueIndex:uidx_username"`
	Password       string    `json:"-" gorm:"type:varchar"` // hashing
	Email          string    `json:"email" gorm:"type:varchar(50);"`
	Role           string    `json:"role" gorm:"type:varchar;"`
	Status         string    `json:"status" gorm:"type:varchar;"`
	ApprovedStatus string    `json:"approved_status" gorm:"type:varchar(30);"`
	CreatedBy      string    `json:"-" gorm:"type:varchar;"`
	UpdatedBy      string    `json:"-" gorm:"type:varchar;"`
	DeletedBy      string    `json:"-" gorm:"type:varchar;"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"-"`
}
