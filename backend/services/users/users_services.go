package service_users

import (
	"backend/db"
	domain_users "backend/domain/users"
	dto_users "backend/dtos/users"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func IsEmailInUse(email string) bool {
	var user domain_users.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return true
	}
	return false
}

func CreateUser(request dto_users.CreateUserRequest) error {
	user := domain_users.User{
		Username: request.Username,
		Password: request.Password,
		Tipo:     "Alumno",
		Email:    request.Email,
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	if IsEmailInUse(user.Email) {
		return errors.New("Su email ya fue usado por otro usuario, ingrese un email diferente")
	}

	user.Password = hashedPassword

	if err := db.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func Login(request dto_users.LoginRequest) (dto_users.LoginResponse, error) {
	var user domain_users.User
	if err := db.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto_users.LoginResponse{}, errors.New("user not found")
		}
		return dto_users.LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return dto_users.LoginResponse{}, errors.New("invalid password")
	}

	token, err := GenerateJWT(request.Username)
	if err != nil {
		return dto_users.LoginResponse{}, err
	}

	return dto_users.LoginResponse{
		Token:  token,
		Tipo:   user.Tipo, // Añadimos el tipo de usuario
		UserID: user.ID,   // Agregar el ID del usuario
	}, nil
}

func ValidateTokenAndGetUserID(tokenString string) (uint, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	var user domain_users.User
	if err := db.DB.Where("username = ?", claims.Username).First(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}
