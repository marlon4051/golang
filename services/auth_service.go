package services

import (
	"errors"
	"task-api/models"
	"task-api/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo  *repositories.UserRepository
	JwtSecret string
}

func (s *AuthService) Login(email, password string) (string, string, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", "", err
	}
	if user == nil {
		return "", "", errors.New("user not found")
	}

	// Compare pass
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("wrong password")
	}

	// Generate JWT
	token, err := s.generateJWT(user.ID)
	if err != nil {
		return "", "", err
	}

	return token, user.UserName, nil
}

func (s *AuthService) Register(user *models.User) error {
	// verify if the user exist
	existingUser, err := s.UserRepo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// create user in db
	return s.UserRepo.CreateUser(user)
}

// Generar JWT
func (s *AuthService) generateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24 hours to expire

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.JwtSecret))
}
