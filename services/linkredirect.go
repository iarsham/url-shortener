package services

import (
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/models"
)

type linkRedirectService struct {
	db *gorm.DB
}

func LinkRedirectRepositoryImpl(db *gorm.DB) domain.LinkRedirectRepository {
	return &linkRedirectService{
		db: db,
	}
}

func (l *linkRedirectService) CheckLinkExists(key string) (models.Link, bool) {
	var dbLink models.Link
	err := l.db.Where("keyword = ?", key).First(&dbLink).Error
	if err != nil {
		return dbLink, false
	}
	return dbLink, true
}
