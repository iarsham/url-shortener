package domain

import (
	"github.com/iarsham/url-shortener/models"
)

type ShortLinkRepository interface {
	CreateShortLink(link *models.Link) error
	RandomShortURL() (string, string)
	CheckLinkExists(url, userID string) bool
	CheckShortLinkExists(shortURL string) bool
	CustomShortURL(key string) string
	ValidateLink(url string) bool
}
