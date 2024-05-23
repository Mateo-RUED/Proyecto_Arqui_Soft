package users

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`  // Store hashed passwords
    Tipo     string `gorm:"type:enum('Administrador', 'Alumno');not null"`
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string `json:"token"`
}

type TokenRequest struct {
    Username string `json:"username"`
}

