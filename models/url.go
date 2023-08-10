package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	LongUrl  string `json:"longurl" gorm:"not null"`
	ShortUrl string `json:"shorturl" gorm:"type:varchar(75);not null"`
	Keyword      string `json:"keyword" gorm:"type:varchar(75);not null"`
	UserID   string `json:"user_id" gorm:"type:varchar(75);not null;index"`
	User     User   `gorm:"forigenkey:UserID"`
}
