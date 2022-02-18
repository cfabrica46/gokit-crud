package service

import (
	_ "github.com/lib/pq"
)

type serviceInterface interface {
	SignUp(string, string, string) (string, error)
	SignIn(string, string) (string, error)
}

type service struct {
	dbHost, dbPort, tokenHost, tokenPort string
}

func GetService(dbHost, dbPort, tokenHost, tokenPort string) *service {
	return &service{dbHost, dbPort, tokenHost, tokenPort}
}

/* func (s service) SignUp(username, password, email string) (token string, err error) {
	//convert password to Sha256
	password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	url := fmt.Sprintf("%s:%s/user/username_password", s.dbHost, s.dbPort)

	dataJSON, err := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		username,
		password,
	})
	if err != nil {
		return
	}

	http.Post(url, "application/json", bytes.NewBuffer(nil))
	// goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	return
}
*/
/* func (serviceInterface) SignIn(username, password string) (token string, err error) {
	return
} */
