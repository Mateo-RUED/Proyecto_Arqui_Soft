package dto_users

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Tipo  string `json:"tipo"` // Añadimos el campo Tipo
	UserID uint   `json:"user_id"` // ID del usuario
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type TokenRequest struct {
	Username string `json:"username"`
}
