package services

import (
	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/models"
)

type verifyUserService struct {
	userRepository domain.UserRepository
}

func VerifyUserServiceImpl(userRepo domain.UserRepository) domain.VerifyUserRepository {
	return &verifyUserService{
		userRepository: userRepo,
	}
}

func (v *verifyUserService) GetUserByEmail(email string) (models.User, error) {
	return v.userRepository.GetUserByEmail(email)
}

func (v *verifyUserService) GetUserFromCache(key string) (models.User, error) {
	return v.userRepository.GetUserFromCache(key)
}

func (v *verifyUserService) Save(user *models.User) error {
	return v.userRepository.Save(user)
}

func (v *verifyUserService) CheckUserStatus(user *models.User) bool {
	if user.IsActive {
		return true
	} else {
		return false
	}
}
