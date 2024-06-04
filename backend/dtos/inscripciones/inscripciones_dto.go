package dto_inscripciones

type ListCoursesByUserRequest struct {
    UsuarioID uint `json:"usuario_id"`
}

type ListCoursesByUserResponse struct {
    Courses []CourseInfo `json:"courses"`
}

type CourseInfo struct {
    ID          uint    `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Precio       float64 `json:"precio"`
    Category    string  `json:"category"`
}
