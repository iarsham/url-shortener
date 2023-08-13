package domain

import "github.com/iarsham/url-shortener/models"

type UserRepository interface {
	Create(user *models.User) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(userID string) (models.User, error)
	Delete(user *models.User) error
	EncryptPassword(password string) (string, error)
	VerifyPassword(hashPass, plainPass string) (bool, error)
	Save(user *models.User) error
	SendVerifyEmail(email string) error
	GetUserFromCache(key string) (models.User, error)
}
