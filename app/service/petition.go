package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
)

/* type MyRequest interface {
	Do()
	Result() (any, error)
}

type MRGetIDByUsername struct {
	client     httpClient
	request    dbapp.UsernameRequest
	response   dbapp.IDErrorResponse
	err        error
	url        string
	methodHTTP string
}

func NewMRGetIDByUsername(client httpClient, url, username string) *MRGetIDByUsername {
	return &MRGetIDByUsername{
		client: client,
		request: dbapp.UsernameRequest{
			Username: username,
		},
		url:        url + "/id/username",
		methodHTTP: http.MethodPost,
	}
}

func (mr *MRGetIDByUsername) Do() {
	bodyJSON, err := json.Marshal(mr.request)
	if err != nil {
		mr.err = fmt.Errorf("error to make petition: %w", err)

		return
	}

	ctx, ctxCancel := context.WithTimeout(context.TODO(), time.Minute)
	defer ctxCancel()

	req, err := http.NewRequestWithContext(ctx, mr.methodHTTP, mr.url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		mr.err = fmt.Errorf("error to make petition: %w", err)

		return
	}

	resp, err := mr.client.Do(req)
	if err != nil {
		mr.err = fmt.Errorf("error to make petition: %w", err)

		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&mr.response)
	if err != nil {
		mr.err = fmt.Errorf("error to make petition: %w", err)

		return
	}
}

func (mr *MRGetIDByUsername) Result() (any, error) {
	return mr.response, mr.err
}

func DoRequest(mr MyRequest) (result any, err error) {
	mr.Do()

	result, err = mr.Result()
	if err != nil {
		err = fmt.Errorf("error to Request: %w", err)

		return nil, err
	}

	return result, nil
} */

// ---

func DoFunc[responseEntity dbapp.IDErrorResponse](client httpClient, body any, url, methodHTTP string, response responseEntity) (resp2 *responseEntity, err error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		err = fmt.Errorf("error to make petition: %w", err)

		return nil, err
	}

	ctx, ctxCancel := context.WithTimeout(context.TODO(), time.Minute)
	defer ctxCancel()

	req, err := http.NewRequestWithContext(ctx, methodHTTP, url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		err = fmt.Errorf("error to make petition: %w", err)

		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("error to make petition: %w", err)

		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode request: %w", err)
	}

	return &response, nil
}

/* func DecodeRequest[req dbapp.IDErrorResponse](request req) httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (any, error) {
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, fmt.Errorf("failed to decode request: %w", err)
		}

		return request, nil
	}
} */
