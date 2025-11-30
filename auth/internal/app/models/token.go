package models

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       uint32 `json:"user_id"`
}
