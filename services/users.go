package services

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/helpers"
	"github.com/iarsham/url-shortener/models"
)

type userRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func UserRepositoryImpl(db *gorm.DB, rdb *redis.Client) domain.UserRepository {
	return &userRepository{db: db, rdb: rdb}
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

func (u *userRepository) VerifyPassword(hashPass, plainPass string) (bool, error) {
	return helpers.VerifyHash(hashPass, plainPass)
}

func (u *userRepository) Save(user *models.User) error {
	err := u.db.Save(&user).Error
	return err
}

func (u *userRepository) SendVerifyEmail(email string) error {
	key, err := helpers.SendVerify(email)
	ttl, _ := strconv.Atoi(os.Getenv("EXPIRE_VERIFY_LINK_Minutes"))
	u.rdb.Set(context.Background(), key, email, time.Minute*time.Duration(ttl))
	return err
}

func (u *userRepository) GetUserFromCache(key string) (models.User, error) {
	var dbUser models.User
	value, err := u.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return dbUser, err
	}
	dbUser, err = u.GetUserByEmail(value)
	if err != nil {
		return dbUser, err
	}
	u.rdb.Del(context.Background(), key)
	return dbUser, nil
}

func (u *userRepository) GetUserWithLinks(userID string) (models.User, error) {
	var dbUser models.User
	err := u.db.Preload("Links").Where("id = ?",userID).First(&dbUser).Error
	return dbUser, err
}
