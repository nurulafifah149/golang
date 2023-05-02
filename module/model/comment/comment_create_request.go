package comment

type CommentCreateRequest struct {
	PhotoId int    `json:"photo_id"`
	Message string `json:"message"`
}
