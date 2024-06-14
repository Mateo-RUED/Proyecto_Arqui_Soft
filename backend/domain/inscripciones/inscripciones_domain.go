package domain_inscripciones

import "time"

type Inscripcion struct {
    ID     uint      `gorm:"primaryKey"`
    Usuario_id uint   `gorm:"not null"`
    Curso_id   uint   `gorm:"not null"`
    Fecha     time.Time `gorm:"not null"`
}
// TableName retorna el nombre de la tabla en la base de datos.
func (Inscripcion) TableName() string {
    return "inscripciones"
}

