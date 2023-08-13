package services

import (
	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/helpers"
	"github.com/iarsham/url-shortener/models"
)

type loginService struct {
	userRepository domain.UserRepository
}

func LoginRepositoryImpl(userRepository domain.UserRepository) domain.LoginRepository {
	return &loginService{
		userRepository: userRepository,
	}
}

func (l *loginService) GetUserByEmail(email string) (models.User, error) {
	return l.userRepository.GetUserByEmail(email)
}

func (l *loginService) CreateAccessToken(userID, email string) string {
	return helpers.GenerateJWT(userID, email)
}

func (u *loginService) VerifyPassword(hashPass, plainPass string) (bool, error) {
	return u.userRepository.VerifyPassword(hashPass, plainPass)
}
