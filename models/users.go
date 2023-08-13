package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:varchar(50);primary_key;default:(UUID())"`
	Email    string    `json:"email" gorm:"type:varchar(150);index;unique;not null" form:"email" binding:"required"`
	IsActive bool      `json:"is_active" gorm:"not null;default:false"`
	Password string    `json:"-" gorm:"type:varchar(150);not null" form:"password" binding:"required"`
	links    []Link    `gorm:"forigenkey:UserID"`
}
