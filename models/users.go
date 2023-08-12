package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:varchar(50);primary_key;default:(UUID())"`
	UserInfo BaseUser  `gorm:"embedded"`
	gorm.Model
}
