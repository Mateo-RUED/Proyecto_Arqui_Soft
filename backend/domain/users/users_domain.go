package domain_users

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`  // Store hashed passwords
    Tipo     string `gorm:"type:enum('Administrador', 'Alumno');not null"`
    Email    string `gorm:"unique;not null"`

}

