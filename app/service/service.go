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
	var response dbapp.InsertUserResponse
	//convert password to Sha256
	password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	// make request
	url := fmt.Sprintf("%s:%s/user", s.dbHost, s.dbPort)
	dataReq, err := json.Marshal(dbapp.InsertUserRequest{username, password, email})
	if err != nil {
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(dataReq))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// recive response
	dataResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataReq, &response)
	if err != nil {
		return
	}

	return
} */

/* func (serviceInterface) SignIn(username, password string) (token string, err error) {
	url := fmt.Sprintf("%s:%s/user/username_password", s.dbHost, s.dbPort)
	return
} */
