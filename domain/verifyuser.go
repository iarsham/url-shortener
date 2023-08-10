package domain

import "github.com/iarsham/url-shortener/models"

type VerifyUserRepository interface {
	GetUserByEmail(email string) (models.User, error)
	GetUserFromCache(key string) (models.User, error)
	CheckUserStatus(user *models.User) bool
	Save(user *models.User) error
}
