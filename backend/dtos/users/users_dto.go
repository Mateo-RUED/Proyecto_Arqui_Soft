package dto_users

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string `json:"token"`
}

type CreateUserRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Tipo     string `json:"tipo"`
    Email    string `json:"email"`
}

type TokenRequest struct {
    Username string `json:"username"`
}

