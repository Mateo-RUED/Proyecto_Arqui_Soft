package domain_inscripciones

import "time"

type Inscripcion struct {
    ID     uint      `gorm:"primaryKey"`
    UsuarioID uint   `gorm:"not null"`
    CursoID   uint   `gorm:"not null"`
    Fecha     time.Time `gorm:"not null"`
}
// TableName retorna el nombre de la tabla en la base de datos.
func (Inscripcion) TableName() string {
    return "inscripciones"
}

