package db

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "log"
    "backend/dao"
	"fmt"
	"time"
)

var DB *gorm.DB

func Init() {
    
    dsn := "root:root@tcp(127.0.0.1:3306)/venta_de_cursos?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
}
func SelectCoursesWithFilter(query string) ([]dao.Course, error) {
	var courses []dao.Course
	result := DB.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func SelectCourseByID(id int64) (dao.Course, error) {
	var course dao.Course
	result := DB.First(&course, id)
	if result.Error != nil {
		return dao.Course{}, fmt.Errorf("not found course with ID: %d", id)
	}
	return course, nil
}
/* 
func InsertSubscription(userID int64, courseID int64) error {
	var subscription dao.Subscription
	result := DB.Where("user_id = ? AND course_id = ?", userID, courseID).First(&subscription)
	if result.Error == nil {
		return fmt.Errorf("user %d is already subscribed to course %d", userID, courseID)
	}

	subscription = dao.Subscription{
		UserID:       userID,
		CourseID:     courseID,
		CreationDate: time.Now().UTC(),
		LastUpdated:  time.Now().UTC(),
	}

	result = DB.Create(&subscription)
	if result.Error != nil {
		return result.Error
	}

	return nil
} */

