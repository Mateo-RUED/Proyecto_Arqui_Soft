package services_archivos

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"backend/db"
	domain_archivos "backend/domain/archivos"
	dtos_archivos "backend/dtos/archivos"
)

const uploadDir = "backend/uploads" // Directorio raíz para archivos

// ArchivoService proporciona métodos para la gestión de archivos
type ArchivoService struct{}

// UploadFile guarda un archivo en el sistema y lo registra en la base de datos
func (s *ArchivoService) UploadFile(cursoID uint, fileName string, fileContent []byte) error {
	// Convertir cursoID a string correctamente
	cursoDir := filepath.Join(uploadDir, fmt.Sprintf("%d", cursoID))

	// Crear la carpeta del curso si no existe
	if _, err := os.Stat(cursoDir); os.IsNotExist(err) {
		if err := os.MkdirAll(cursoDir, 0755); err != nil {
			return fmt.Errorf("error al crear el directorio del curso %s: %w", cursoDir, err)
		}
	}

	// Construir la ruta completa del archivo
	filePath := filepath.Join(cursoDir, fileName)

	// Guardar el archivo en el sistema
	if err := os.WriteFile(filePath, fileContent, 0644); err != nil {
		return fmt.Errorf("error al guardar el archivo %s: %w", filePath, err)
	}

	// Registrar el archivo en la base de datos
	archivo := domain_archivos.Archivo{
		CursoID:     cursoID,
		Nombre:      fileName,
		RutaArchivo: filePath,
	}
	if err := db.DB.Create(&archivo).Error; err != nil {
		return fmt.Errorf("error al registrar el archivo en la base de datos: %w", err)
	}

	return nil
}

// GetFilesByCourse devuelve una lista de archivos para un curso específico
func (s *ArchivoService) GetFilesByCourse(cursoID uint) ([]dtos_archivos.ArchivoResponse, error) {
	var archivos []domain_archivos.Archivo
	if err := db.DB.Where("curso_id = ?", cursoID).Find(&archivos).Error; err != nil {
		return nil, errors.New("error al obtener los archivos")
	}

	// Crear la lista de respuesta con URLs para descargar
	var response []dtos_archivos.ArchivoResponse
	for _, archivo := range archivos {
		response = append(response, dtos_archivos.ArchivoResponse{
			Nombre: archivo.Nombre,
			URL:    fmt.Sprintf("/archivos/%d/descargar/%s", cursoID, archivo.Nombre), // Generar la URL para descarga
		})
	}

	return response, nil
}

// DownloadFile devuelve la ruta de un archivo específico
func (s *ArchivoService) DownloadFile(cursoID uint, fileName string) (string, error) {
	var archivo domain_archivos.Archivo
	// Corregir la consulta para usar "nombre_archivo"
	if err := db.DB.Where("curso_id = ? AND nombre_archivo = ?", cursoID, fileName).First(&archivo).Error; err != nil {
		return "", errors.New("archivo no encontrado")
	}

	// Verificar si el archivo existe físicamente
	if _, err := os.Stat(archivo.RutaArchivo); os.IsNotExist(err) {
		return "", errors.New("el archivo no existe en el servidor")
	}

	return archivo.RutaArchivo, nil
}
