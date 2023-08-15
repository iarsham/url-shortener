package services

import (
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

func (s *shortLinkService) CreateShortLink(link *models.Link) error {
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

func (s *shortLinkService) CheckShortLinkExists(shortURL string) bool {
	var dbLink models.Link
	s.db.Where("short_url = ?", shortURL).First(&dbLink)
	return dbLink.ID != 0
}

func (s *shortLinkService) RandomShortURL() (string, string) {
	return helpers.MakeShortURL()
}

func (s *shortLinkService) CustomShortURL(key string) string {
	return helpers.CustomShortURL(key)
}
