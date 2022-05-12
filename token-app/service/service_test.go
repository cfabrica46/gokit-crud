package service_test

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/cfabrica46/gokit-crud/token-app/service"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	idTest       int    = 1
	usernameTest string = "username"
	emailTest    string = "email@email.com"
	secretTest   string = "secret"
	tokenTest    string = "token"

	errRedisClosed string = "redis: client is closed"
)

func TestGenerateToken(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                string
		inUsername, inEmail string
		outToken, outErr    string
		inSecret            []byte
		inID                int
	}{
		{
			name:       "NoError",
			inID:       idTest,
			inUsername: usernameTest,
			inEmail:    emailTest,
			inSecret:   []byte(secretTest),
			outToken:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.",
			outErr:     "",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result string

			mr, err := miniredis.Run()
			if err != nil {
				assert.Error(t, err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			result = svc.GenerateToken(tt.inID, tt.inUsername, tt.inEmail, tt.inSecret)

			assert.Contains(t, result, tt.outToken)
		})
	}
}

func TestExtractToken(t *testing.T) {
	t.Parallel()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       idTest,
		"username": usernameTest,
		"email":    emailTest,
		"uuid":     uuid.NewString(),
	})

	tokenSigned, err := token.SignedString([]byte(secretTest))
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name                          string
		inToken                       string
		outUsername, outEmail, outErr string
		inSecret                      []byte
		outID                         int
	}{
		{
			name:        "NoError",
			inToken:     tokenSigned,
			inSecret:    []byte(secretTest),
			outID:       idTest,
			outUsername: usernameTest,
			outEmail:    emailTest,
			outErr:      "",
		},
		{
			name:        "NotValidToken",
			inToken:     "",
			inSecret:    nil,
			outID:       0,
			outUsername: "",
			outEmail:    "",
			outErr:      "token contains an invalid number of segments",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultID int
			var resultUsername, resultEmail, resultErr string

			mr, err := miniredis.Run()
			if err != nil {
				assert.Error(t, err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			resultID, resultUsername, resultEmail, err = svc.ExtractToken(tt.inToken, tt.inSecret)
			if err != nil {
				resultErr = err.Error()
			}

			if tt.name == "NoError" {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			assert.Equal(t, tt.outID, resultID, "they should be equal")
			assert.Equal(t, tt.outUsername, resultUsername, "they should be equal")
			assert.Equal(t, tt.outEmail, resultEmail, "they should be equal")
		})
	}
}

func TestSetToken(t *testing.T) {
	t.Parallel()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       idTest,
		"username": usernameTest,
		"email":    emailTest,
		"uuid":     uuid.NewString(),
	})

	tokenSigned, _ := token.SignedString([]byte(secretTest))

	for _, tt := range []struct {
		name   string
		in     string
		outErr string
	}{
		{
			name:   "NoError",
			in:     tokenSigned,
			outErr: "",
		},
		{
			name:   "ErrorRedisClose",
			in:     "",
			outErr: "redis: client is closed",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			mr, err := miniredis.Run()
			if err != nil {
				assert.Error(t, err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			// Generate Conflict.
			if tt.outErr == "redis: client is closed" {
				svc.DB.Close()
			}

			err = svc.SetToken(tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			if tt.name == "NoError" {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}
		})
	}
}

func TestDeleteToken(t *testing.T) {
	t.Parallel()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       idTest,
		"username": usernameTest,
		"email":    emailTest,
		"uuid":     uuid.NewString(),
	})

	tokenSigned, _ := token.SignedString([]byte(secretTest))

	for _, tt := range []struct {
		name   string
		in     string
		outErr string
	}{
		{
			name:   "NoError",
			in:     tokenSigned,
			outErr: "",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			mr, err := miniredis.Run()
			if err != nil {
				assert.Error(t, err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			err = svc.DeleteToken(tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			if tt.name == "NoError" {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}
		})
	}
}

func TestCheckToken(t *testing.T) {
	t.Parallel()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       idTest,
		"username": usernameTest,
		"email":    emailTest,
		"uuid":     uuid.NewString(),
	})

	tokenSigned, _ := token.SignedString([]byte(secretTest))

	for _, tt := range []struct {
		name     string
		in       string
		outErr   string
		outCheck bool
	}{
		{
			name:     "NoError",
			in:       tokenSigned,
			outCheck: true,
			outErr:   "",
		},
		{
			name:     "NoError",
			in:       "",
			outCheck: false,
			outErr:   "",
		},
		{
			name:     "ErrorRedisClose",
			in:       "",
			outCheck: false,
			outErr:   "redis: client is closed",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultCheck bool
			var resultErr string

			mr, err := miniredis.Run()
			if err != nil {
				assert.Error(t, err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			// insert.
			if tt.in != "" {
				err = svc.SetToken(tt.in)
				if err != nil {
					assert.Error(t, err)
				}
			}

			// Generate Conflict.
			if tt.outErr == "redis: client is closed" {
				svc.DB.Close()
			}

			resultCheck, err = svc.CheckToken(tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			if tt.name == "NoError" {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			assert.Equal(t, tt.outCheck, resultCheck, "they should be equal")
		})
	}
}

func TestKeyFunc(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                string
		inUsername, inEmail string
		outErr              string
		inSecret            []byte
		outSecret           []byte
		inID                int
	}{
		{
			name:       "Error",
			inSecret:   []byte(secretTest),
			inID:       idTest,
			inUsername: usernameTest,
			inEmail:    emailTest,
			outSecret:  []byte(nil),
			outErr:     service.ErrUnexpectedSigningMethod.Error(),
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result []byte
			var resultErr string

			kf := service.KeyFunc(tt.inSecret)

			// generateToken.
			token := jwt.NewWithClaims(jwt.SigningMethodPS256, jwt.MapClaims{
				"id":       tt.inID,
				"username": tt.inUsername,
				"email":    tt.inEmail,
				"uuid":     uuid.NewString(),
			})

			res, err := kf(token)
			if err != nil {
				resultErr = err.Error()
			}

			if tt.name == "NoError" {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			result, ok := res.([]byte)
			if resultErr == "" {
				if !ok {
					assert.Fail(t, "response is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outSecret, result)
		})
	}
}
