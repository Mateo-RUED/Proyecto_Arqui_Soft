package domain_courses

type Course struct {
    ID          uint    `gorm:"primaryKey"`
    Name        string  `gorm:"not null"`
    Description string  `gorm:"not null"`
    Category     string  `gorm:"not null"`
    Requisitos string `gorm:"not null"`
    Duracion string `gorm:"not null"`
    Imagen_url string `gorm:"not null"`
}
