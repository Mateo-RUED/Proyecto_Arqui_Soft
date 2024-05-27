package dtos

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
}

type TokenRequest struct {
    Username string `json:"username"`
}
