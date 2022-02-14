package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type serviceInterface interface {
	GenerateToken(int, string, string, []byte) (string, error)
	ExtractData(string, []byte) (int, string, string, error)
	SetToken(string) error
	DeleteToken(string) error
	CheckToken(string) (bool, error)
}

type service struct {
	db   *redis.Client
	once sync.Once
}

func GetService() *service {
	return &service{}
}

func (s *service) OpenDB() (err error) {
	s.once.Do(func() {
		options := &redis.Options{
			Addr:     RedisHost + ":" + RedisPort,
			Password: "",
			DB:       0,
		}
		s.db = redis.NewClient(options)
	})
	return
}

func (service) GenerateToken(id int, username, email string, secret []byte) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
		"uuid":     uuid.NewString(),
	})

	token, err = t.SignedString(secret)
	return
}

func (s service) ExtractData(token string, secret []byte) (id int, username, email string, err error) {
	t, err := jwt.Parse(token, keyFunc(secret))
	if err != nil {
		return
	}

	claims := t.Claims.(jwt.MapClaims)
	idAux := claims["id"].(float64)
	id = int(idAux)
	username = claims["username"].(string)
	email = claims["email"].(string)
	return
}

func (s *service) SetToken(token string) (err error) {
	err = s.db.Set(token, true, time.Minute*10).Err()
	if err != nil {
		return
	}
	return
}

func (s *service) DeleteToken(token string) error {
	return s.db.Del(token).Err()
}

func (s service) CheckToken(token string) (check bool, err error) {
	result, err := s.db.Get(token).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			err = nil
			return
		}
		return
	}

	if result == "1" {
		check = true
	}
	return
}

func keyFunc(secret []byte) func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	}
}
