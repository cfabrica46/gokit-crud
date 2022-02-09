package service

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestDecodeGetAllUsersRequest(t *testing.T) {
	for i, tt := range []struct {
		in       *http.Request
		out      getAllUsersRequest
		outError error
	}{
		{&http.Request{}, getAllUsersRequest{}, nil},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := decodeGetAllUsersRequest(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeGetUserByIDRequest(t *testing.T) {
	id := 1
	url := fmt.Sprintf("loclahost:8080/user/%d", id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Error(err)
	}

	for i, tt := range []struct {
		in       *http.Request
		out      getUserByIDRequest
		outError error
	}{
		{req, getUserByIDRequest{}, nil},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := decodeGetUserByIDRequest(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

/* func TestDecodeGetUserByUsernameAndPasswordRequest(t *testing.T) {
	id := 1
	url := fmt.Sprintf("loclahost:8080/user/%d", id)

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}
	for i, tt := range []struct {
		in       *http.Request
		out      getUserByIDRequest
		outError error
	}{
		{req, getUserByIDRequest{}, nil},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := decodeGetUserByIDRequest(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
} */

/* func decodeGetIDByUsernameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getIDByUsernameRequest

	username := mux.Vars(r)["username"]
	request.Username = username

	return request, nil
} */

/* func decodeInsertUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request insertUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
} */

/* func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
*/
/* func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
} */
