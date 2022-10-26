package response

import "time"

type CommentCreateResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentUserGetAllResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CommentPhotoGetAllResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type CommentGetAllResponse struct {
	ID        uint                       `json:"id"`
	Message   string                     `json:"message"`
	PhotoID   uint                       `json:"photo_id"`
	UserID    uint                       `json:"user_id"`
	UpdatedAt time.Time                  `json:"updated_at"`
	CreatedAt time.Time                  `json:"created_at"`
	User      CommentUserGetAllResponse  `json:"user"`
	Photo     CommentPhotoGetAllResponse `json:"photo"`
}

type CommentUpdateResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentDeleteResponse struct {
	Message string `json:"message"`
}
