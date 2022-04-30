package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const (
	lifeOfToken int = 10
)

var ErrUnexpectedSigningMethod = errors.New("unexpected signing method")

type serviceInterface interface {
	GenerateToken(int, string, string, []byte) string
	ExtractToken(string, []byte) (int, string, string, error)
	SetToken(string) error
	DeleteToken(string) error
	CheckToken(string) (bool, error)
}

// Service ...
type Service struct {
	DB *redis.Client
}

// GetService ...
func GetService(db *redis.Client) *Service {
	return &Service{db}
}

// GenerateToken ...
func (Service) GenerateToken(id int, username, email string, secret []byte) (token string) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
		"uuid":     uuid.NewString(),
	})

	token, _ = t.SignedString(secret)

	return token
}

// ExtractToken ...
func (Service) ExtractToken(token string, secret []byte) (id int, username, email string, err error) {
	t, err := jwt.Parse(token, KeyFunc(secret))
	if err != nil {
		return 0, "", "", fmt.Errorf("error to extract token: %w", err)
	}

	claims, _ := t.Claims.(jwt.MapClaims)
	idAux, _ := claims["id"].(float64)
	id = int(idAux)
	username, _ = claims["username"].(string)
	email, _ = claims["email"].(string)

	return id, username, email, nil
}

// SetToken ...
func (s *Service) SetToken(token string) (err error) {
	err = s.DB.Set(token, true, time.Minute*time.Duration(lifeOfToken)).Err()
	if err != nil {
		return fmt.Errorf("error to set token: %w", err)
	}

	return nil
}

// DeleteToken ...
func (s *Service) DeleteToken(token string) (err error) {
	if err := s.DB.Del(token).Err(); err != nil {
		return fmt.Errorf("failed to delete token: %w", err)
	}

	return nil
}

// CheckToken ...
func (s Service) CheckToken(token string) (check bool, err error) {
	result, err := s.DB.Get(token).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}

		return false, fmt.Errorf("error to get token: %w", err)
	}

	if result == "1" {
		check = true
	}

	return check, nil
}

func KeyFunc(secret []byte) func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}

		return secret, nil
	}
}
