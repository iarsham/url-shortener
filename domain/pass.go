package domain

import "github.com/iarsham/url-shortener/models"

type PasswordRepository interface {
	GetUserByID(userID string) (models.User, error)
	EncryptPassword(password string) (string, error)
	Save(user *models.User) error
}
