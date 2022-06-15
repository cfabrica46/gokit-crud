package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	tokenapp "github.com/cfabrica46/gokit-crud/token-app/service"
)

type MyResponse interface {
	dbapp.UserErrorResponse |
		dbapp.IDErrorResponse |
		dbapp.ErrorResponse |
		dbapp.RowsErrorResponse |
		tokenapp.Token |
		tokenapp.IDUsernameEmailErrResponse |
		tokenapp.ErrorResponse |
		tokenapp.CheckErrResponse
}

func RequestFunc[responseEntity MyResponse](client HttpClient, body any, url, methodHTTP string, response *responseEntity) (err error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		err = fmt.Errorf("error to make petition: %w", err)

		return err
	}

	ctx, ctxCancel := context.WithTimeout(context.TODO(), time.Minute)
	defer ctxCancel()

	req, err := http.NewRequestWithContext(ctx, methodHTTP, url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		err = fmt.Errorf("error to make petition: %w", err)

		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("error to make petition: %w", err)

		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("failed to decode request: %w", err)
	}

	return nil
}
