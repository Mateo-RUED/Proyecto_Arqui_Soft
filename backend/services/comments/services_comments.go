package services_comments

import (
	"backend/db"
	domain_comments "backend/domain/comments"
	dto_comments "backend/dtos/comments"
	"errors"
	"time"
)
func AddComment(request dto_comments.AddCommentRequest) error {
	// Validar el contenido del comentario
	if request.Content == "" {
		return errors.New("el contenido del comentario no puede estar vacío")
	}
	// Crear el comentario
	comment := domain_comments.Comment{
		CourseID:  request.CourseID,
		UserID:    request.UserID,
		Content:   request.Content,
		CreatedAt: time.Now(),
	}
	// Guardar el comentario en la base de datos
	if err := db.DB.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}
func GetCommentsByCourse(courseID uint) ([]dto_comments.CommentResponse, error) {
	var comments []domain_comments.Comment
	// Obtener comentarios del curso
	if err := db.DB.Where("course_id = ?", courseID).Find(&comments).Error; err != nil {
		return nil, err
	}
	// Mapear a la respuesta
	var response []dto_comments.CommentResponse
	for _, comment := range comments {
		response = append(response, dto_comments.CommentResponse{
			ID:        comment.ID,
			UserID:    comment.UserID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return response, nil
}