package service

import (
	"database/sql"
	"log"
)

type serviceInterface interface {
	GetAllUsers() ([]User, error)
	GetUserByID(int) (User, error)
	GetUserByUsernameAndPassword(string, string) (User, error)
	GetIDByUsername(string) (int, error)
	InsertUser(string, string, string) error
	DeleteUser(int) (int, error)
}

// Service ...
type Service struct {
	db *sql.DB

	// Data for DB
	host, port, user, password, dbName, sslmode, driver string
}

//GetService ...
func GetService(db *sql.DB) *Service {
	return &Service{db: db}
}

//GetAllUsers ...
func (s Service) GetAllUsers() (users []User, err error) {
	log.SetFlags(log.Lshortfile)
	rows, err := s.db.Query("SELECT id, username, email FROM users")
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

//GetUserByID ...
func (s Service) GetUserByID(id int) (user User, err error) {
	row := s.db.QueryRow("SELECT id, username, password, email FROM users WHERE id = $1", id)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//GetUserByUsernameAndPassword ...
func (s Service) GetUserByUsernameAndPassword(username, password string) (user User, err error) {
	row := s.db.QueryRow("SELECT id, username, password, email FROM users WHERE username = $1 AND password = $2", username, password)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//GetIDByUsername ...
func (s Service) GetIDByUsername(username string) (id int, err error) {
	row := s.db.QueryRow("SELECT id FROM users WHERE username = $1", username)

	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//InsertUser ...
func (s *Service) InsertUser(username, password, email string) (err error) {
	_, err = s.db.Exec("INSERT INTO users(username, password, email) VALUES ($1,$2,$3)", username, password, email)
	if err != nil {
		return
	}

	return
}

//DeleteUser ...
func (s *Service) DeleteUser(id int) (rowsAffected int, err error) {
	r, err := s.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return
	}

	count, _ := r.RowsAffected()
	rowsAffected = int(count)
	return
}
