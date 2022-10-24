package request

type PhotoCreateRequest struct {
	Title    string `binding:"required" json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `binding:"required" json:"photo_url"`
}

type PhotoUpdateRequest struct {
	Title    string `binding:"required" json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `binding:"required" json:"photo_url"`
}
