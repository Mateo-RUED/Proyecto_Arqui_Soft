package users

import (
    "time"
    "errors"
    "github.com/dgrijalva/jwt-go"
    "backend/dtos/users"
    "backend/domain/users"
    "backend/db"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
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

// HashPassword hashes a password using bcrypt.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func ValidateUserType(userType string) error {
    if userType != "Administrador" && userType != "Alumno" {
        return errors.New("invalid user type")
    }
    return nil
}

// CreateUser creates a new user in the database.
func CreateUser(user users.User) error {
    if err := ValidateUserType(user.Tipo); err != nil {
        return err
    }
    
    hashedPassword, err := HashPassword(user.Password)
    if err != nil {
        return err
    }
    
    userDTO := dtos.CreateUserRequest{
        Username: user.Username,
        Password: hashedPassword,
        Tipo:     user.Tipo,
    }
    
    return db.DB.Create(&userDTO).Error
}

// Login verifies the user credentials and generates a JWT for a login request.
func Login(request users.LoginRequest) (users.LoginResponse, error) {
    var user users.User
    if err := db.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return users.LoginResponse{}, errors.New("user not found")
        }
        return users.LoginResponse{}, err
    }

    // Compare the provided password with the hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
        return users.LoginResponse{}, errors.New("invalid password")
    }

    token, err := GenerateJWT(request.Username)
    if err != nil {
        return users.LoginResponse{}, err
    }

    return users.LoginResponse{
        Token: token,
    }, nil
}




