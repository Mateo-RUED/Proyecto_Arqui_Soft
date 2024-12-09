package domain_archivos

// Archivo representa un archivo relacionado con un curso
type Archivo struct {
    ID          uint   `gorm:"primaryKey"`
    CursoID     uint   `gorm:"not null"`
    Nombre      string `gorm:"column:nombre_archivo;not null"`
    RutaArchivo string `gorm:"column:ruta_archivo;not null"`
}
