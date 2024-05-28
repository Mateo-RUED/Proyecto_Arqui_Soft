package dto

type CreateCourseDTO struct {
    Name        string `json:"name" validate:"required"`
    Description string `json:"description" validate:"required"`
}

type UpdateCourseDTO struct {
    Name        string `json:"name" validate:"required"`
    Description string `json:"description" validate:"required"`
}