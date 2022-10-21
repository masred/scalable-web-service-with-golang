package domain

type SocialMedia struct {
	Name           string `gorm:"not null" validate:"required" json:"name"`
	SocialMediaUrl string `gorm:"not null" validate:"required" json:"social_media_url"`
	UserID         string `json:"user_id"`
}
