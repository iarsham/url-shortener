package entity

type LinkRequest struct {
	URL   string `json:"url" form:"url" binding:"required"`
	Param string `json:"param" form:"param"  binding:"omitempty,min=7,max=7"`
}
