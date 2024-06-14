package services_inscripciones

import (
    "backend/domain/courses"
    "backend/domain/users"
    "backend/domain/inscripciones"
    "backend/dtos/inscripciones"
    "backend/db"
    "errors"
    "time"
)

// InscribirUsuario inscribe a un usuario en un curso.
func InscribirUsuario(request dto_inscripciones.InscribirUsuarioRequest) error {
     // Verificar si el curso existe
    var course domain_courses.Course

    if err := db.DB.Where("id = ?", request.CourseID).First(&course).Error; err != nil {
        return errors.New("curso no encontrado: " + err.Error())
    }

    // Verificar si el usuario existe
    var user domain_users.User
    if err := db.DB.Where("id = ?", request.UserID).First(&user).Error; err != nil {
        return errors.New("usuario no encontrado: " + err.Error())
    } 

    // Crear la inscripción
    inscripcion := domain_inscripciones.Inscripcion{

    Usuario_id: request.UserID,
    Curso_id:   request.CourseID,
    Fecha:    time.Now(),

    }

    // Guardar la inscripción en la base de datos
    if err := db.DB.Create(&inscripcion).Error; err != nil {
        return err
    }

    return nil
}

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
        if err := db.DB.First(&course, inscripcion.Curso_id).Error; err != nil {
            return dto_inscripciones.ListCoursesByUserResponse{}, err
        }
        courses = append(courses, dto_inscripciones.CourseInfo{
            ID:          course.ID,
            Name:        course.Name,
            Description: course.Description,
            Category:    course.Category,
            Requisitos: course.Requisitos,
            Duracion: course.Duracion,
            Imagen_url: course.Imagen_url,
        })
    }

    return dto_inscripciones.ListCoursesByUserResponse{Courses: courses}, nil
} 

