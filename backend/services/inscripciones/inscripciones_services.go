package services_inscripciones

import (
    "backend/domain/courses"
    "backend/domain/inscripciones"
    "backend/dtos/inscripciones"
    "backend/db"
    "errors" // Importamos el paquete errors
)

func ListCoursesByUser(usuarioID uint) (dto_inscripciones.ListCoursesByUserResponse, error) {
    var inscripciones []domain_inscripciones.Inscripcion
    if err := db.DB.Where("usuario_id = ?", usuarioID).Find(&inscripciones).Error; err != nil {
        return dto_inscripciones.ListCoursesByUserResponse{}, err
    }

    var courses []dto_inscripciones.CourseInfo
    if len(inscripciones) == 0 {
        return dto_inscripciones.ListCoursesByUserResponse{}, errors.New("no courses found for the user") // Usamos errors aquí
    }

    for _, inscripcion := range inscripciones {
        var course domain_courses.Course
        if err := db.DB.First(&course, inscripcion.CursoID).Error; err != nil {
            return dto_inscripciones.ListCoursesByUserResponse{}, err
        }
        courses = append(courses, dto_inscripciones.CourseInfo{
            ID:          course.ID,
            Name:        course.Name,
            Description: course.Description,
            Precio:      course.Precio,
            Category:    course.Category,
        })
    }

    return dto_inscripciones.ListCoursesByUserResponse{Courses: courses}, nil
}

