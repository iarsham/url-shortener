package domain

import "github.com/iarsham/url-shortener/models"

type ShortLinkRepository interface {
	CreateShortLink(link *models.Link) error
	CheckLinkExists(url ,userID string ) bool
	ValidateLink(url string) bool
}
