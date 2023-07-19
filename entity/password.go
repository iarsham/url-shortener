package entity

type PasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required,min=8"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=8"`
}
