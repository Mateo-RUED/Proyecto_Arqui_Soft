package domain

type Course struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"type:varchar(100)"`
    Description string `gorm:"type:varchar(255)"`
}
type Courses []Course