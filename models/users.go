package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/helpers"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:varchar(50);primary_key;default:(UUID())"`
	UserInfo BaseUser  `gorm:"embedded"`
	gorm.Model
}

type UserLogin struct {
	BaseUser
}

func (u *User) HashUserPassword(tx *gorm.DB) error {
	_, err := helpers.Hash(u.UserInfo.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) VerifyUserPassword(p string) (bool, error) {
	return helpers.VerifyHash(u.UserInfo.Password, p)
}
