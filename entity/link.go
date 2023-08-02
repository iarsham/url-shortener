package entity

type LinkRequest struct {
	URL string `json:"url" form:"url" binding:"required"`
}
