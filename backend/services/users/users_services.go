package users

import (
    "time"
    "github.com/dgrijalva/jwt-go"
    usersDomain "backend/domain/users"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// GenerateJWT generates a JWT for the given username.
func GenerateJWT(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// Login generates a JWT for a login request.
func Login(request usersDomain.LoginRequest) usersDomain.LoginResponse {
    token, err := GenerateJWT(request.Username)
    if err != nil {
        // Handle the error according to your error handling strategy
        return usersDomain.LoginResponse{
            Token: "",
        }
    }
    return usersDomain.LoginResponse{
        Token: token,
    }
}

