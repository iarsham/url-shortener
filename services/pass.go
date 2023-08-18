package services

import (
	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/models"
)

type passwordService struct {
	userRepository domain.UserRepository
}

func PasswordServiceImpl(userRepo domain.UserRepository) domain.PasswordRepository {
	return &passwordService{
		userRepository: userRepo,
	}
}

func (p *passwordService) GetUserByID(id string) (models.User, error) {
	return p.userRepository.GetUserByID(id)
}

func (p *passwordService) EncryptPassword(password string) string{
	return p.userRepository.EncryptPassword(password)
}

func (p *passwordService) Save(user *models.User) error {
	return p.userRepository.Save(user)
}

func (u *passwordService) VerifyPassword(hashPass, plainPass string) (bool, error) {
	return u.userRepository.VerifyPassword(hashPass, plainPass)
}
