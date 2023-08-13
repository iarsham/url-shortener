package domain

import "github.com/iarsham/url-shortener/models"

type LoginRepository interface {
	GetUserByEmail(email string) (models.User, error)
	VerifyPassword(hashPass, plainPass string) (bool, error)
	CreateAccessToken(userID, email string) string 
}
