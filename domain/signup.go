package domain

import "github.com/iarsham/url-shortener/models"

type SignUpRepository interface {
	Create(user *models.User) error
	GetUserByEmail(email string) (models.User, error)
	CreateAccessToken(userID, email string) string
	EncryptPassword(password string) (string, error)
	SendVerifyEmail(email string) error
}
