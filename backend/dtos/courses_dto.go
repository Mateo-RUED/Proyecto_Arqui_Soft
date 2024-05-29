package dtos

type CreateCourseDTO struct {
    Name        string `json:"name" binding:"required"`
    Description string `json:"description" binding:"required"`
}

type UpdateCourseDTO struct {
    Name        string `json:"name"`
    Description string `json:"description"`
}
