package courses

type CreateCourseDTO struct {
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description" binding:"required"`
    Category    string  `json:"category" binding:"required"`
    Price       float64 `json:"price" binding:"required"`
}


/* type UpdateCourseDTO struct {
    ID          int64  `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}  */

// dto/course_dto.go
/* package dtos

type CreateCourseDTO struct {
    Name        string `json:"name" binding:"required"`
    Description string `json:"description"`
}

type UpdateCourseDTO struct {
    Name        string `json:"name" binding:"required"`
    Description string `json:"description"`
}

type CourseDetailDto struct {
    Id          uint   `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

type CoursesDetailDto []CourseDetailDto */

