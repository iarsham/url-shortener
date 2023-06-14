package models


type BaseUser struct {
	Email    string `json:"email" gorm:"type:varchar(150);index;unique;not null" form:"email" binding:"required"`
	Password string `json:"password" gorm:"type:varchar(150);not null" form:"password" binding:"required"`
}
