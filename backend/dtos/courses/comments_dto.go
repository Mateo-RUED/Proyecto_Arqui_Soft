package dto_courses

type AddCommentRequest struct {
	CourseID uint   `json:"course_id"`
	UserID   uint   `json:"user_id"`
	Content  string `json:"content"`
}

type CommentResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
