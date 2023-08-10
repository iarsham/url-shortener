package entity


type Authenticate struct {
	Email string `json:"email" form:"email" binding:"required" example:"james@yahoo.com"`
	Password string `json:"password" form:"password" binding:"required,min=8" example:"password!@#123"`
}
