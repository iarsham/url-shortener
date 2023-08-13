package helpers

import (
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/models"
)

func RemoveExpiredUrls(db *gorm.DB) {
	minutes, _ := strconv.Atoi(os.Getenv("EXPIRE_URLS_MINUTES"))
	duration := time.Now().Add(-time.Duration(minutes) * time.Minute)
	err := db.Where("created_at < ?", duration).Delete(&models.Link{}).Error
	if err != nil {
		log.Fatal(err)
	}
}
