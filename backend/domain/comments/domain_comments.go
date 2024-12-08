package domain_comments

import "time"
type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	CourseID  uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}
// TableName especifica el nombre de la tabla en la base de datos
func (Comment) TableName() string {
	return "comments"
}