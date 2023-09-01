package services

import (
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/configs"
	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/models"
)

type linkRedirectService struct {
	db *gorm.DB
	*configs.CustomLogger
}

func LinkRedirectRepositoryImpl(db *gorm.DB, lg *configs.CustomLogger) domain.LinkRedirectRepository {
	return &linkRedirectService{
		db:           db,
		CustomLogger: lg,
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

func (l *linkRedirectService) Save(link *models.Link) error {
	err := l.db.Save(&link).Error
	if err != nil {
		l.Logger.Fatal(err.Error())
		return err
	}
	l.Logger.Printf("link (%s) received a new view", link.ShortUrl)
	return err
}
