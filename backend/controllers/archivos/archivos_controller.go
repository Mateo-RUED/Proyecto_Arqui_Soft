package controller_archivos

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const uploadDir = "backend/uploads" // Directorio raíz para archivos

// UploadFile permite a los administradores subir un archivo relacionado con un curso.
func UploadFile(c *gin.Context) {
	cursoID := c.Param("cursoID")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo obtener el archivo"})
		return
	}

	// Crear una subcarpeta para el curso si no existe
	cursoDir := filepath.Join(uploadDir, cursoID)
	if _, err := os.Stat(cursoDir); os.IsNotExist(err) {
		if mkdirErr := os.MkdirAll(cursoDir, os.ModePerm); mkdirErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al crear el directorio"})
			return
		}
	}

	// Guardar el archivo
	filePath := filepath.Join(cursoDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al guardar el archivo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Archivo subido exitosamente", "file": file.Filename})
}

// GetFilesByCourse lista los archivos subidos para un curso específico
func GetFilesByCourse(c *gin.Context) {
	cursoID := c.Param("cursoID")
	cursoDir := filepath.Join(uploadDir, cursoID)

	// Verificar si la carpeta existe
	files, err := os.ReadDir(cursoDir)
	if os.IsNotExist(err) {
		c.JSON(http.StatusOK, gin.H{"files": []string{}})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al leer los archivos"})
		return
	}

	// Crear la lista de archivos
	var fileList []string
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}

	c.JSON(http.StatusOK, gin.H{"files": fileList})
}

// DownloadFile permite descargar un archivo específico
func DownloadFile(c *gin.Context) {
	filename := c.Param("filename")
	cursoID := c.Param("cursoID")

	filePath := filepath.Join(uploadDir, cursoID, filename)

	// Verificar si el archivo existe
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Archivo no encontrado"})
		return
	}

	// Establecer cabeceras para descarga
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.File(filePath)
}
