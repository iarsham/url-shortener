package entity

type LoginSignUpOkResponse struct {
	Response string `json:"response" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImI0Y2MzODE4NDk2QG15bWFpbHkubG9sIiwiZXhwIjoxNjkwOTIxNjMxLCJ1c2VyX2lkIjoiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIn0.Vs2BXM2Z6hr4zqLLWe08FrpKhDfRpnaFhu4TKB5Spb4"`
}

type User404Responsse struct {
	Response string `json:"response" example:"user not found"`
}

type PasswordIncorrectResponsse struct {
	Response string `json:"response" example:"password is incorrect"`
}

type DataBodyResponse struct {
	Response string `json:"response" example:"body properties required"`
}

type UserExistResponse struct {
	Response string `json:"response" example:"user with this email already exists"`
}

type DBErrorResponse struct {
	Response string `json:"response" example:"failed to save data to db"`
}

type LinkExpireResponse struct {
	Response string `json:"response" example:"link is invalid or expired"`
}

type VerifyOKResponse struct {
	Response string `json:"response" example:"user verified successfully"`
}

type AlreadyVerifiedResponse struct {
	Response string `json:"response" example:"user already verified"`
}

type PasswordOkResponse struct {
	Response string `json:"response" example:"password changed successfully"`
}

type NewPasswordEqualResponse struct {
	Response string `json:"response" example:"new passwords must be equal"`
}

type IncorrectCurrentPasswordResponse struct {
	Response string `json:"response" example:"current password is incorrect"`
}

type ShortLinkOkResponse struct {
	Response string `json:"response" example:"http://domain/yjMl1vl"`
}

type ShortLinkDBErrorResponse struct {
	Response string `json:"response" example:"cant short long url"`
}

type ShortLinkValidateResponse struct {
	Response string `json:"response" example:"url is invalid"`
}
type ShortLinkExistsResponse struct {
	Response string `json:"response" example:"url already exists"`
}

type ShortLinkNotExistsResponse struct {
	Response string `json:"response" example:"url not found"`
}
