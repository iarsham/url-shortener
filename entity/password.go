package entity

type PasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required,min=8" example:"James!123"`
	Password        string `json:"password" binding:"required,min=8" example:"1qaz2wsx"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=8" example:"1qaz2wsx"`
}
