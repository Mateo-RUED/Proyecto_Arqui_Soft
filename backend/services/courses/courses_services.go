package services

import (
    "lms/database"
    "lms/domain"
    "lms/dto"
    "gorm.io/gorm"
)

type CourseService struct {
    DB *gorm.DB
}

func NewCourseService(db *gorm.DB) *CourseService {
    return &CourseService{
        DB: db,
    }
}
func (s *CourseService) CreateCourse(courseDTO dto.CreateCourseDTO) (*domain.Course, error) {
    course := domain.Course{
        Name:        courseDTO.Name,
        Description: courseDTO.Description,
    }
    if err := s.DB.Create(&course).Error; err != nil {
        return nil, err
    }
    return &course, nil
}

func (s *CourseService) UpdateCourse(id uint, courseDTO dto.UpdateCourseDTO) (*domain.Course, error) {
    var course domain.Course
    if err := s.DB.First(&course, id).Error; err != nil {
        return nil, err
    }
    course.Name = courseDTO.Name
    course.Description = courseDTO.Description
    if err := s.DB.Save(&course).Error; err != nil {
        return nil, err
    }
    return &course, nil
}

func (s *CourseService) DeleteCourse(id uint) error {
    var course domain.Course
    if err := s.DB.First(&course, id).Error; err != nil {
        return err
    }
    return s.DB.Delete(&course).Error
}