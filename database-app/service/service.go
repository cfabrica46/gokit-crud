package service

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

type serviceDBInterface interface {
	OpenDB() error
	InsertUser() error
	DeleteUserByUsername() error
	// GetAllUsers() (users []models.User, err error)
}

type serviceDB struct {
	db   *sql.DB
	once sync.Once
}

func getServiceDB() serviceDB {
	return serviceDB{once: sync.Once{}}
}

func (s *serviceDB) OpenDB(dbDriver, psqlInfo string) (err error) {
	s.once.Do(func() {
		s.db, err = sql.Open(dbDriver, psqlInfo)
		if err != nil {
			return
		}
		err = s.db.Ping()
		if err != nil {
			return
		}
	})
	return
}

func (s *serviceDB) InsertUser(username, password, email string) (err error) {
	stmt, err := s.db.Prepare("INSERT INTO users(username,password,email) VALUES ($1,$2,$3)")
	if err != nil {
		return
	}

	_, err = stmt.Exec(username, password, email)
	if err != nil {
		return
	}
	return
}

func (s *serviceDB) DeleteUserByUsername(username string) (rowsAffected int, err error) {
	stmt, err := s.db.Prepare("DELETE FROM users WHERE username = $1")
	if err != nil {
		return
	}

	r, _ := stmt.Exec(username)
	count, _ := r.RowsAffected()
	rowsAffected = int(count)
	return
}

/* func (serviceDB) GetAllUsers() (users []models.User, err error) {
	return
} */
