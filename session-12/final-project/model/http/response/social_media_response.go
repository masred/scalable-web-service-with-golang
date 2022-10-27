package response

import "time"

type SocialMediaCreateResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type SocialMediaUserGetAllResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type SocialMediaGetAllResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
	User           SocialMediaUserGetAllResponse
}

type SocialMediaUpdateResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaDeleteResponse struct {
	Message string `json:"message"`
}
