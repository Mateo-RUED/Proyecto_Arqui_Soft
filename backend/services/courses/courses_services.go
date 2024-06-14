package services_courses

import (
    "backend/domain/courses"
    "backend/dtos/courses"
    "backend/db"
    "errors"
    "gorm.io/gorm"

)

// CreateCourse crea un nuevo curso en la base de datos.
func CreateCourse(request dto_courses.CreateCourseRequest) error {
    course := domain_courses.Course{
        Name:        request.Name,
        Description: request.Description,
        Category:    request.Category,
        Requisitos: request.Requisitos,
        Duracion: request.Duracion,
        Imagen_url: request.Imagen_url,
    }

    if CourseExists(course.Name) {
        return errors.New("el curso ya existe, ingrese otro nombre de curso")
    }

    if err := ValidateCategory(course.Category); err != nil {
        return err
    }

    // Guardar el curso en la base de datos
    if err := db.DB.Create(&course).Error; err != nil {
        return err
    }
    
    return nil
}

func ValidateCategory(category string) error {
    validCategories := []string{"Programación", "Análisis de datos", "Base de datos", "Desarrollo web"}
    for _, validCategory := range validCategories {
        if category == validCategory {
            return nil
        }
    }
    return errors.New("Categoría inválida. Ingrese alguna de las siguientes:Programación, Análisis de datos, Base de datos, Desarrollo web")
}

func CourseExists(courseName string) bool {
    var course domain_courses.Course
    if err := db.DB.Where("name = ?", courseName).First(&course).Error; err != nil {
        return false // No se encontró el curso, por lo que no existe
    }
    return true // Se encontró el curso, por lo que ya existe
}

func DeleteCurso(courseName string) error {
    var course domain_courses.Course
    result := db.DB.Where("name = ?", courseName).First(&course)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return errors.New("No existe curso con ese nombre")
        }
        return result.Error
    }

    if err := db.DB.Delete(&course).Error; err != nil {
        return err
    }

    return nil
}

func EditCourse(editRequest dto_courses.EditCourseRequest) error {
    // Verificar si el curso existe en la base de datos
    var existingCourse domain_courses.Course
    result := db.DB.First(&existingCourse, editRequest.ID)
    if result.Error != nil {
        return errors.New("Curso no encontrado")
    }

    // Actualizar los campos del curso con los valores proporcionados
    existingCourse.Name = editRequest.Name
    existingCourse.Description = editRequest.Description
    existingCourse.Category = editRequest.Category
    existingCourse.Requisitos = editRequest.Requisitos
    existingCourse.Duracion = editRequest.Duracion
    existingCourse.Imagen_url = editRequest.Imagen_url


    if err := ValidateCategory(existingCourse.Category); err != nil {
        return err
    }

    // Guardar los cambios en la base de datos
    if err := db.DB.Save(&existingCourse).Error; err != nil {
        return err
    }

    return nil
}

func GetCourseByID(id uint) (dto_courses.GetCourseByIDResponse, error) {
    var course domain_courses.Course
    if err := db.DB.First(&course, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return dto_courses.GetCourseByIDResponse{}, errors.New("course not found")
        }
        return dto_courses.GetCourseByIDResponse{}, err
    }

    response := dto_courses.GetCourseByIDResponse{
        ID:          course.ID,
        Name:        course.Name,
        Description: course.Description,
        Category:    course.Category,
        Requisitos: course.Requisitos,
        Duracion: course.Duracion,
        Imagen_url: course.Imagen_url,
    }

    return response, nil
}

func DeleteCourseByID(id uint) error {
    if err := db.DB.Delete(&domain_courses.Course{}, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("course not found")
        }
        return err
    }
    return nil
}


func GetCoursesByCategory(category string) ([]dto_courses.GetCourseByCategoryResponse, error) {
    var courses []domain_courses.Course
    if err := db.DB.Where("category = ?", category).Find(&courses).Error; err != nil {
        return nil, err
    }
    var response []dto_courses.GetCourseByCategoryResponse
    for _, course := range courses {
        response = append(response, dto_courses.GetCourseByCategoryResponse{
            ID:          course.ID,
            Name:        course.Name,
            Description: course.Description,
            Category:    course.Category,
            Requisitos:  course.Requisitos,
            Duracion:    course.Duracion,
            Imagen_url:  course.Imagen_url,
        })
    }

    return response, nil
} 

// GetAllCourses devuelve todos los cursos de la base de datos.
func GetAllCourses() ([]dto_courses.GetAllCoursesResponse, error) {
    var courses []domain_courses.Course
    if err := db.DB.Find(&courses).Error; err != nil {
        return nil, err
    }

    var response []dto_courses.GetAllCoursesResponse
    for _, course := range courses {
        response = append(response, dto_courses.GetAllCoursesResponse{
            ID:          course.ID,
            Name:        course.Name,
            Description: course.Description,
            Category:    course.Category,
            Requisitos:  course.Requisitos,
            Duracion:    course.Duracion,
            Imagen_url:  course.Imagen_url,
        })
    }

    return response, nil
}
