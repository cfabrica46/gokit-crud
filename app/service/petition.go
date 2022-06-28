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

type HTTPComponents struct {
	url, method string
}

func NewHTTPComponents(url, method string) HTTPComponents {
	return HTTPComponents{url: url, method: method}
}

func RequestFunc[responseEntity MyResponse](
	client HTTPClient,
	body any,
	httpComponents HTTPComponents,
	response *responseEntity,
) (err error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		err = fmt.Errorf("error to make petition: %w", err)

		return err
	}

	ctx, ctxCancel := context.WithTimeout(context.TODO(), time.Minute)
	defer ctxCancel()

	req, err := http.NewRequestWithContext(ctx, httpComponents.method, httpComponents.url, bytes.NewBuffer(bodyJSON))
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

func RequestFuncWithoutBody(
	client HTTPClient,
	httpComponents HTTPComponents,
	response *dbapp.UsersErrorResponse,
) (err error) {
	ctx, ctxCancel := context.WithTimeout(context.TODO(), time.Minute)
	defer ctxCancel()

	req, err := http.NewRequestWithContext(ctx, httpComponents.method, httpComponents.url, bytes.NewBuffer(nil))
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
