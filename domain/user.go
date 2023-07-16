package domain

import "github.com/iarsham/url-shortener/models"

type UserRepository interface {
	Create(user *models.User) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(userID string) (models.User, error)
	Delete(user *models.User) error
}
