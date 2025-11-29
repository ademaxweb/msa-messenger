package models

type User struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}
