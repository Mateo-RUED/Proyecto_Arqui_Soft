package dto_courses

type CreateCourseRequest struct {
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Precio       float64 `json:"price"`
    Category     string  `json:"category"`
}

type DeleteCourseRequest struct {
    Name string `json:"nombre"`
}

type EditCourseRequest struct {
    ID          uint   `json:"id"`  // ID del curso a editar
    Name      string `json:"nombre"`
    Description string `json:"descripcion"`
    Precio      float64 `json:"precio"`
    Category   string `json:"categoria"`
}

type GetCourseByIDRequest struct {
    ID          uint   `json:"id"`
}

type GetCourseByIDResponse struct {
    ID          uint   `json:"id"`
    Name        string `json:"nombre"`
    Description string `json:"descripcion"`
    Precio      float64 `json:"precio"`
    Category    string `json:"categoria"`
}
