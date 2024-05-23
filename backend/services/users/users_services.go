package users

import (
    "time"
    "errors"
    "github.com/dgrijalva/jwt-go"
    usersDomain "backend/domain/users"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "backend/db"
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
func CreateUser(user usersDomain.User) error {

    if err := ValidateUserType(user.Tipo); err != nil {
        return err
    }
    
    hashedPassword, err := HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword
    return db.DB.Create(&user).Error
}

// Login verifies the user credentials and generates a JWT for a login request.
func Login(request usersDomain.LoginRequest) (usersDomain.LoginResponse, error) {
    var user usersDomain.User
    if err := db.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return usersDomain.LoginResponse{}, errors.New("user not found")
        }
        return usersDomain.LoginResponse{}, err
    }

    // Compare the provided password with the hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
        return usersDomain.LoginResponse{}, errors.New("invalid password")
    }

    token, err := GenerateJWT(request.Username)
    if err != nil {
        return usersDomain.LoginResponse{}, err
    }

    return usersDomain.LoginResponse{
        Token: token,
    }, nil
}


