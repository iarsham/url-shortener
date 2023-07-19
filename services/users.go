package services

import (
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/helpers"
	"github.com/iarsham/url-shortener/models"
)

type userRepository struct {
	db *gorm.DB
}

func UserRepositoryImpl(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(user *models.User) error {
	err := u.db.Create(&user).Error
	return err
}

func (u *userRepository) GetUserByEmail(email string) (models.User, error) {
	var dbUser models.User
	err := u.db.Where("email = ?", email).First(&dbUser).Error
	return dbUser, err
}

func (u *userRepository) GetUserByID(userID string) (models.User, error) {
	var dbUser models.User
	err := u.db.Where("id = ?", userID).First(&dbUser).Error
	return dbUser, err
}

func (u *userRepository) Delete(user *models.User) error {
	err := u.db.Delete(&user).Error
	return err
}

func (u *userRepository) EncryptPassword(password string) (string, error) {
	return helpers.Hash(password)
}

func (u *userRepository) Save(user *models.User) error {
	err := u.db.Save(&user).Error
	return err
}
