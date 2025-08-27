package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"workshop2/model"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var JwtSecret = []byte("supersecretkey")

func HashPassword(pw string) string {
	h := sha256.New()
	h.Write([]byte(pw))
	return hex.EncodeToString(h.Sum(nil))
}

func CheckPassword(input, hashed string) bool {
	return HashPassword(input) == hashed
}

func RegisterUser(db *gorm.DB, req *model.User) error {
	if req.Email == "" || req.Password == "" {
		return ErrMissingFields
	}
	var exists model.User
	db.Where("email = ?", req.Email).First(&exists)
	if exists.ID != 0 {
		return ErrEmailExists
	}
	req.Password = HashPassword(req.Password)
	return db.Create(req).Error
}

func LoginUser(db *gorm.DB, email, password string) (string, error) {
	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 || !CheckPassword(password, user.Password) {
		return "", ErrInvalidCredentials
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	t, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetUserByID(db *gorm.DB, id int) (*model.User, error) {
	var user model.User
	db.First(&user, id)
	if user.ID == 0 {
		return nil, ErrUserNotFound
	}
	user.Password = ""
	return &user, nil
}

var ErrMissingFields = errors.New("missing email or password")
var ErrEmailExists = errors.New("email already registered")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUserNotFound = errors.New("user not found")
