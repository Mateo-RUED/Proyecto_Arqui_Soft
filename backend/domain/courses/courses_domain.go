package domain

type Course struct {
    Id          uint   `gorm:"primaryKey;autoIncrement"`
    Name        string `gorm:"type:varchar(255);not null"`
    Description string `gorm:"type:varchar(1000);not null"`
}

type Courses []Course