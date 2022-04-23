package service

import (
	"database/sql"
	"errors"
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
}

// GetService ...
func GetService(db *sql.DB) *Service {
	return &Service{db: db}
}

// GetAllUsers ...
func (s Service) GetAllUsers() (users []User, err error) {
	rows, err := s.db.Query("SELECT id, username, email FROM users")
	if err != nil {
		return
	}
	defer rows.Close()

	// TODO: verify scan erros type values
	// TODO: rows.Err not rows
	for rows.Next() {
		var userBeta User

		err = rows.Scan(&userBeta.ID, &userBeta.Username, &userBeta.Email)
		if err != nil {
			return
		}

		users = append(users, userBeta)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

// GetUserByID ...
func (s Service) GetUserByID(id int) (user User, err error) {
	row := s.db.QueryRow("SELECT id, username, password, email FROM users WHERE id = $1", id)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}

		return
	}

	return
}

// GetUserByUsernameAndPassword ...
func (s Service) GetUserByUsernameAndPassword(username, password string) (user User, err error) {
	row := s.db.QueryRow(
		"SELECT id, username, password, email FROM users WHERE username = $1 AND password = $2",
		username,
		password,
	)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}

		return
	}

	return
}

// GetIDByUsername ...
func (s Service) GetIDByUsername(username string) (id int, err error) {
	row := s.db.QueryRow("SELECT id FROM users WHERE username = $1", username)

	err = row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}

		return
	}

	return
}

// InsertUser ...
func (s *Service) InsertUser(username, password, email string) (err error) {
	_, err = s.db.Exec(
		"INSERT INTO users(username, password, email) VALUES ($1,$2,$3)",
		username,
		password,
		email,
	)
	if err != nil {
		return
	}

	return
}

// DeleteUser ...
func (s *Service) DeleteUser(id int) (rowsAffected int, err error) {
	r, err := s.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return
	}

	count, _ := r.RowsAffected()
	rowsAffected = int(count)

	return
}
