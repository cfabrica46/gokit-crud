package service

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

type serviceInterface interface {
	GetAllUsers() ([]User, error)
	GetUserByID(int) (User, error)
	GetUserByUsernameAndPassword(string, string) (User, error)
	GetIDByUsername(string) (int, error)
	InsertUser(string, string, string) error
	DeleteUser(string, string, string) (int, error)
}

type service struct {
	db   *sql.DB
	once sync.Once

	// Data for DB
	host, port, user, password, dbName, sslmode, driver string
}

func GetService(host, port, user, password, dbName, sslmode, driver string) *service {
	return &service{once: sync.Once{}, host: host, port: port, user: user, password: password, dbName: dbName, sslmode: sslmode, driver: driver}
}

func (s *service) OpenDB() (err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", s.host, s.port, s.user, s.password, s.dbName, s.sslmode)

	s.once.Do(func() {
		s.db, err = sql.Open(s.driver, psqlInfo)
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

func (s service) GetAllUsers() (users []User, err error) {
	rows, err := s.db.Query("SELECT users.id,users.username,users.email FROM users")
	if err != nil {
		return
	}

	for rows.Next() {
		var userBeta User
		rows.Scan(&userBeta.ID, &userBeta.Username, &userBeta.Email)
		users = append(users, userBeta)
	}
	return
}

func (s service) GetUserByID(id int) (user User, err error) {
	row := s.db.QueryRow("SELECT users.id,users.username,users.password,users.email FROM users WHERE users.id = $1", id)

	var userBeta User
	err = row.Scan(&userBeta.ID, &userBeta.Username, &userBeta.Password, &userBeta.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	user = userBeta
	return
}

func (s service) GetUserByUsernameAndPassword(username, password string) (user User, err error) {
	row := s.db.QueryRow("SELECT users.id, users.email FROM users WHERE users.username = $1 AND users.password = $2", username, password)

	var userBeta User

	err = row.Scan(&userBeta.ID, &userBeta.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	user = userBeta
	user.Username = username
	user.Password = password
	return
}

func (s service) GetIDByUsername(username string) (id int, err error) {
	row := s.db.QueryRow("SELECT users.id FROM users WHERE users.username = $1", username)

	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

func (s *service) InsertUser(username, password, email string) (err error) {
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

func (s *service) DeleteUser(username, password, email string) (rowsAffected int, err error) {
	stmt, err := s.db.Prepare("DELETE FROM users WHERE username = $1 AND password = $2 AND email = $3")
	if err != nil {
		return
	}

	r, _ := stmt.Exec(username, password, email)
	count, _ := r.RowsAffected()
	rowsAffected = int(count)
	return
}
