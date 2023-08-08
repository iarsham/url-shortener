package domain

import (
	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/models"
)

type LinkRedirectRepository interface {
	CheckLinkExists(url string) (models.Link, bool)
	GetCurrentHost(ctx *gin.Context) string
}
