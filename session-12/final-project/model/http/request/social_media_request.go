package request

type SocialMediaCreateRequest struct {
	Name           string `binding:"required" json:"name"`
	SocialMediaUrl string `binding:"required" json:"social_media_url"`
}

type SocialMediaUpdateRequest struct {
	Name           string `binding:"required" json:"name"`
	SocialMediaUrl string `binding:"required" json:"social_media_url"`
}
