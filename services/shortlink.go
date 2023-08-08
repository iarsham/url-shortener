package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/helpers"
	"github.com/iarsham/url-shortener/models"
)

type shortLinkService struct {
	db *gorm.DB
}

func ShortLinkRepositoryImpl(db *gorm.DB) domain.ShortLinkRepository {
	return &shortLinkService{
		db: db,
	}
}

func (s *shortLinkService) CreateShortLink(link *models.Link, ctx *gin.Context) error {
	return s.db.Create(&link).Error
}

func (s *shortLinkService) ValidateLink(url string) bool {
	return helpers.IsValidURL(url)
}

func (s *shortLinkService) CheckLinkExists(url, userID string) bool {
	var dbLink models.Link
	s.db.Where("long_url = ? AND user_id = ?", url, userID).First(&dbLink)
	return dbLink.ID != 0
}

func (s *shortLinkService) RandomShortURL(ctx *gin.Context) string {
	return helpers.MakeShortURL(ctx)
}

func (s *shortLinkService) GetCurrentHost(ctx *gin.Context) string {
	return helpers.CurrentHost(ctx)
}
