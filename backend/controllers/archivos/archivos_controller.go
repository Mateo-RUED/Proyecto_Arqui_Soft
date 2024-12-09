package controller_archivos

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"backend/services/archivos"
)

var archivoService = services_archivos.ArchivoService{}

// UploadFile permite a los administradores subir un archivo relacionado con un curso.
func UploadFile(c *gin.Context) {
	cursoIDStr := c.Param("cursoID")
	cursoID, err := strconv.Atoi(cursoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Curso ID inválido"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo obtener el archivo"})
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al abrir el archivo"})
		return
	}
	defer fileContent.Close()

	content, err := ioutil.ReadAll(fileContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al leer el archivo"})
		return
	}

	if err := archivoService.UploadFile(uint(cursoID), file.Filename, content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Archivo subido exitosamente"})
}

// GetFilesByCourse lista los archivos subidos para un curso específico
func GetFilesByCourse(c *gin.Context) {
	cursoIDStr := c.Param("cursoID")
	cursoID, err := strconv.Atoi(cursoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Curso ID inválido"})
		return
	}

	files, err := archivoService.GetFilesByCourse(uint(cursoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}

// DownloadFile permite descargar un archivo específico
func DownloadFile(c *gin.Context) {
	cursoIDStr := c.Param("cursoID")
	cursoID, err := strconv.Atoi(cursoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Curso ID inválido"})
		return
	}

	filename := c.Param("filename")
	filePath, err := archivoService.DownloadFile(uint(cursoID), filename)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.File(filePath)
}
