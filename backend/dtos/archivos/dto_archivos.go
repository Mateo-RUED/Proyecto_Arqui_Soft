package dtos_archivos

// ArchivoRequest representa la solicitud para subir un archivo
type ArchivoRequest struct {
	CursoID uint   `json:"curso_id" binding:"required"`
	File    string `json:"file" binding:"required"`
}


// ArchivoResponse representa la respuesta para listar o descargar archivos
type ArchivoResponse struct {
	Nombre string `json:"nombre"`
	URL    string `json:"url"`
}
