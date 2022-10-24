package response

import (
	"time"
)

type PhotoCreateResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoUserGetAllReponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoGetAllResponse struct {
	ID        uint                   `json:"id"`
	Title     string                 `json:"title"`
	Caption   string                 `json:"caption"`
	PhotoUrl  string                 `json:"photo_url"`
	UserID    uint                   `json:"user_id"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	User      PhotoUserGetAllReponse `json:"user"`
}

type PhotoUpdateResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoDeleteResponse struct {
	Message string `json:"message"`
}
