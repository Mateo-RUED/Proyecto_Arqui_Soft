package dto_inscripciones

type InscribirUsuarioRequest struct {
	UserID   uint `json:"usuario_id"`
	CourseID uint `json:"curso_id"`
}

/*
	 type ListCoursesByUserRequest struct {
	    UsuarioID uint `json:"usuario_id"`
	}
*/
type ListCoursesByUserResponse struct {
	Courses []CourseInfo `json:"courses"`
}

type CourseInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Requisitos  string `json:"requisitos"`
	Duracion    string `json:"duracion"`
	Imagen_url  string `json:"imagen_url"`
}
