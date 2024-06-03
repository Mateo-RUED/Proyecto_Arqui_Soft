package domain_courses

type Course struct {
    ID          uint    `gorm:"primaryKey"`
    Name        string  `gorm:"not null"`
    Description string  `gorm:"not null"`
    Precio       float64 `gorm:"not null"`
    Category     string  `gorm:"not null"`
}

type Courses []Course 