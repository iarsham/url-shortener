package domain

import (
	"github.com/iarsham/url-shortener/models"
)

type LinkRedirectRepository interface {
	CheckLinkExists(key string) (models.Link, bool)
}
