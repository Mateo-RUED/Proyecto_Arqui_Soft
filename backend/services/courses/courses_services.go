package services

import (

    "backend/db"
    "backend/domain/courses" 
    "backend/dtos"
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

func (s *CourseService) CreateCourse(courseDTO dtos.CreateCourseDTO) (*domain.Course, error) {
    course := domain.Course{
        Name:        courseDTO.Name,
        Description: courseDTO.Description,
    }
      if err := s.DB.Create(&course).Error; err != nil {
        return nil, err
    }  
    return &course, nil
}

func (s *CourseService) UpdateCourse(id uint, courseDTO dtos.UpdateCourseDTO) (*domain.Course, error) {
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
}   */


// services/course_service.go
package services

import (
    "Proyecto_Arqui_Soft/backend/dtos"
    "Proyecto_Arqui_Soft/backend/domain"
    e "Proyecto_Arqui_Soft/backend/utils/errors"
    "gorm.io/gorm"
)

type courseService struct {
    DB *gorm.DB
}

type courseServiceInterface interface {
    CreateCourse(courseDTO dtos.CreateCourseDTO) (*dtos.CourseDetailDto, e.ApiError)
    UpdateCourse(id uint, courseDTO dtos.UpdateCourseDTO) (*dtos.CourseDetailDto, e.ApiError)
    DeleteCourse(id uint) e.ApiError
    GetCourseById(id uint) (dtos.CourseDetailDto, e.ApiError)
    GetCourses() (dtos.CoursesDetailDto, e.ApiError)
}

var (
    CourseService courseServiceInterface
)

func NewCourseService(db *gorm.DB) *courseService {
    return &courseService{DB: db}
}

func init() {
    CourseService = &courseService{}
}

func (s *courseService) CreateCourse(courseDTO dtos.CreateCourseDTO) (*dtos.CourseDetailDto, e.ApiError) {
    course := domain.Course{
        Name:        courseDTO.Name,
        Description: courseDTO.Description,
    }
    if err := s.DB.Create(&course).Error; err != nil {
        return nil, e.NewInternalServerApiError("failed to create course", err)
    }

    courseDetailDto := dtos.CourseDetailDto{
        Id:          course.ID,
        Name:        course.Name,
        Description: course.Description,
    }
    return &courseDetailDto, nil
}

func (s *courseService) UpdateCourse(id uint, courseDTO dtos.UpdateCourseDTO) (*dtos.CourseDetailDto, e.ApiError) {
    var course domain.Course
    if err := s.DB.First(&course, id).Error; err != nil {
        return nil, e.NewNotFoundApiError("course not found")
    }

    course.Name = courseDTO.Name
    course.Description = courseDTO.Description

    if err := s.DB.Save(&course).Error; err != nil {
        return nil, e.NewInternalServerApiError("failed to update course", err)
    }

    courseDetailDto := dtos.CourseDetailDto{
        Id:          course.ID,
        Name:        course.Name,
        Description: course.Description,
    }
    return &courseDetailDto, nil
}

func (s *courseService) DeleteCourse(id uint) e.ApiError {
    var course domain.Course
    if err := s.DB.First(&course, id).Error; err != nil {
        return e.NewNotFoundApiError("course not found")
    }
    if err := s.DB.Delete(&course).Error; err != nil {
        return e.NewInternalServerApiError("failed to delete course", err)
    }
    return nil
}

func (s *courseService) GetCourseById(id uint) (dtos.CourseDetailDto, e.ApiError) {
    var course domain.Course
    if err := s.DB.First(&course, id).Error; err != nil {
        return dtos.CourseDetailDto{}, e.NewNotFoundApiError("course not found")
    }

    courseDetailDto := dtos.CourseDetailDto{
        Id:          course.ID,
        Name:        course.Name,
        Description: course.Description,
    }
    return courseDetailDto, nil
}

func (s *courseService) GetCourses() (dtos.CoursesDetailDto, e.ApiError) {
    var courses []domain.Course
    if err := s.DB.Find(&courses).Error; err != nil {
        return nil, e.NewInternalServerApiError("failed to get courses", err)
    }

    var coursesDto dtos.CoursesDetailDto
    for _, course := range courses {
        courseDetailDto := dtos.CourseDetailDto{
            Id:          course.ID,
            Name:        course.Name,
            Description: course.Description,
        }
        coursesDto = append(coursesDto, courseDetailDto)
    }

    return coursesDto, nil
} */

