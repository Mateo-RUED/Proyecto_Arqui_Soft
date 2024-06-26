package dto_courses

type CreateCourseRequest struct {
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Category     string  `json:"category"`
    Requisitos string `json: "requisitos"`
    Duracion string `json: "duracion"`
    Imagen_url string `json: "imagen_url"`
}

type DeleteCourseRequest struct {
    Name string `json:"nombre"`
}

type EditCourseRequest struct {
    ID          uint   `json:"id"`  // ID del curso a editar
    Name      string `json:"nombre"`
    Description string `json:"descripcion"`
    Category   string `json:"categoria"`
    Requisitos string `json: "requisitos"`
    Duracion string `json: "duracion"`
    Imagen_url string `json: "imagen_url"`
}

type GetCourseByIDRequest struct {
    ID          uint   `json:"id"`
}

type GetCourseByIDResponse struct {
    ID          uint   `json:"id"`  // ID del curso a editar
    Name      string `json:"nombre"`
    Description string `json:"descripcion"`
    Category   string `json:"categoria"`
    Requisitos string `json: "requisitos"`
    Duracion string `json: "duracion"`
    Imagen_url string `json: "imagen_url"`
}
// -------------------

type GetCourseByCategoryRequest struct {
    Category string `json:"categoria"`
}
type GetCourseByCategoryResponse struct {
    ID          uint   `json:"id"`  
    Name      string `json:"nombre"`
    Description string `json:"descripcion"`
    Category   string `json:"categoria"`
    Requisitos string `json: "requisitos"`
    Duracion string `json: "duracion"`
    Imagen_url string `json: "imagen_url"`

} 


type GetAllCoursesResponse struct {
    ID          uint   `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Category    string `json:"category"`
    Requisitos  string `json:"requisitos"`
    Duracion    string `json:"duracion"`
    Imagen_url  string `json:"imagen_url"`
}
