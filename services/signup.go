package services

import (
	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/helpers"
	"github.com/iarsham/url-shortener/models"
)

type signUpService struct {
	userRepository domain.UserRepository
}

func SignUpRepositoryImpl(userRepository domain.UserRepository) domain.SignUpRepository {
	return &signUpService{
		userRepository: userRepository,
	}
}

func (s *signUpService) Create(user *models.User) error {
	return s.userRepository.Create(user)
}

func (s *signUpService) GetUserByEmail(email string) (models.User, error) {
	return s.userRepository.GetUserByEmail(email)
}

func (s *signUpService) CreateAccessToken(userID, email string) string {
	return helpers.GenerateJWT(userID, email)
}

func (s *signUpService) EncryptPassword(password string) (string, error) {
	return s.userRepository.EncryptPassword(password)
}

func (s *signUpService)SendVerifyEmail(email string) error{
	return s.userRepository.SendVerifyEmail(email)
}