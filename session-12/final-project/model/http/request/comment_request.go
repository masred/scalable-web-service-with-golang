package request

type CommentCreateRequest struct {
	Message string `binding:"required" json:"message"`
	PhotoID uint   `json:"photo_id"`
}

type CommentUpdateRequest struct {
	Message string `binding:"required" json:"message"`
}
