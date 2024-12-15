package auth

import (
	"fmt"
	"go-auth/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Register(user models.User) error
	Validate(email, password string) (bool, error)
	ValidateJWT(tokenString string) (jwt.MapClaims, error)
	GenerateJWT(email string) (string, error)
}

type authservice struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) *authservice {
	return &authservice{
		repo: repo,
	}
}

func (s *authservice) Register(user models.User) error {
	isUserExist := s.repo.IsUserExists(user.Email)

	if isUserExist {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	return s.repo.Register(user)
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func (s *authservice) GenerateJWT(email string) (string, error) {

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	// Define claims
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *authservice) ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims if token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func (s *authservice) Validate(email, password string) (bool, error) {
	return s.repo.Validate(email, password)
}
