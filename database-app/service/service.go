package service

import (
	"database/sql"
	"sync"

	"github.com/cfabrica46/gokit-crud/database-app/models"
	_ "github.com/lib/pq"
)

type serviceDBInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(int) (models.User, error)
	GetUserByUsernameAndPassword(string, string) (models.User, error)
	GetIDByUsername(string) (int, error)
	InsertUser(string, string, string) error
	DeleteUserByUsername(string) (int, error)
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

func (s serviceDB) GetAllUsers() (users []models.User, err error) {
	rows, err := s.db.Query("SELECT users.id,users.username,users.email FROM users")
	if err != nil {
		return
	}

	for rows.Next() {
		var userBeta models.User
		rows.Scan(&userBeta.ID, &userBeta.Username, &userBeta.Email)
		users = append(users, userBeta)
	}
	return
}

func (s serviceDB) GetUserByID(id int) (user models.User, err error) {
	row := s.db.QueryRow("SELECT users.id,users.username,users.password,users.email FROM users WHERE users.id = $1", id)

	var userBeta models.User
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

func (s serviceDB) GetUserByUsernameAndPassword(username, password string) (user models.User, err error) {
	row := s.db.QueryRow("SELECT users.id, users.email FROM users WHERE users.username = $1 AND users.password = $2", username, password)

	var userBeta models.User

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

func (s serviceDB) GetIDByUsername(username string) (id int, err error) {
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
