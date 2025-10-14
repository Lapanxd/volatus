package service

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/lapanxd/volatus-api/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, username, password string) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("cannot hash password")
	}

	user := model.User{
		Username: username,
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, errors.New("cannot create user")
	}

	return &user, nil
}

func AuthenticateUser(db *gorm.DB, username, password string) (*model.User, error) {
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func GenerateJWT(userID uint) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtDuration, err := time.ParseDuration(os.Getenv("JWT_EXPIRES"))
	if err != nil {
		jwtDuration = 72 * time.Hour
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(jwtDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
