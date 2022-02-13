package service

import (
	"sync"

	"github.com/go-redis/redis"
)

type serviceInterface interface {
	GenerateToken(int, string, string) (string, error)
	ExtractData(string) (int, string, string, error)
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

		/* s.db, err = sql.Open(dbDriver, psqlInfo)
		if err != nil {
			return
		}
		err = s.db.Ping()
		if err != nil {
			return
		} */
	})
	return
}

func (s service) GenerateToken(id int, username, email string) (token string, err error) {
	return
}

func (s service) ExtractData(token string) (id int, username, password string, err error) {
	return
}
