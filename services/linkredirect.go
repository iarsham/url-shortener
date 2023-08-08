package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/helpers"
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

func (l *linkRedirectService) CheckLinkExists(url string) (models.Link, bool) {
	var dbLink models.Link
	err := l.db.Where("short_url = ?", url).First(&dbLink).Error
	if err != nil {
		return dbLink, false
	}
	return dbLink, true
}

func (l *linkRedirectService) GetCurrentHost(ctx *gin.Context) string {
	return helpers.CurrentHost(ctx)
}
