package domain

import (
	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/models"
)

type ShortLinkRepository interface {
	CreateShortLink(link *models.Link, ctx *gin.Context) error
	RandomShortURL(ctx *gin.Context) string
	CheckLinkExists(url, userID string) bool
	ValidateLink(url string) bool
}
