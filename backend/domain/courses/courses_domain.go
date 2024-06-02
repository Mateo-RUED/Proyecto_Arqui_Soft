  package domain_courses

  type Course struct {
	  ID          int     `gorm:"primaryKey;autoIncrement"`
	  Name        string  `gorm:"type:varchar(100);not null"`
	  Description string  `gorm:"type:text;not null"`
	  Category    string  `gorm:"type:varchar(100);not null"`
	  Price       float64 `gorm:"type:decimal(10,2);not null"`
  }
type Courses []Course  


/* package domain

type Course struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"not null"`
    Description string
}
 */

// package domain
/* 
package courses

import "time"

type Course struct {
	ID           int64     `json:"id"`            // Course ID
	Title        string    `json:"title"`         // Course title
	Description  string    `json:"description"`   // Course description
	Category     string    `json:"category"`      // Course Category. Allowed values: to be defined
	ImageURL     string    `json:"image_url"`     // Course image URL
	CreationDate time.Time `json:"creation_date"` // Course creation date
	LastUpdated  time.Time `json:"last_updated"`  // Course last updated date
}

type SearchResponse struct {
	Results []Course `json:"results"`
}

type SubscribeRequest struct {
	UserID   int64 `json:"user_id"`
	CourseID int64 `json:"course_id"`
} */