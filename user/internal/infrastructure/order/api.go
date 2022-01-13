package order

import (
	"context"
	"epitech/deliveats/user/internal/conf"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type API struct {
	client  *http.Client
	address string
	apiKey  string
}

func NewAPI(config conf.OrderService) *API {
	client := &http.Client{}
	return &API{client: client, address: config.URL, apiKey: config.APIKey}
}

func (a *API) DeleteAll(ctx context.Context, userID string) error {
	uri := fmt.Sprintf("%s/orders", a.address)

	request, err := http.NewRequest(http.MethodDelete, uri, nil)
	if err != nil {
		return errors.Wrap(err, "unable to build http request")
	}

	request = request.WithContext(ctx)
	request.Header.Set("api-key", a.apiKey)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", userID))

	resp, err := a.client.Do(request)
	if err != nil {
		return errors.Wrap(err, "unable to handle the request")
	}
	switch resp.StatusCode {
	case http.StatusNotFound, http.StatusOK:
		return nil
	default:
		return fmt.Errorf("DELETE order API did not respond OK. HTTP code: %d", resp.StatusCode)
	}
}
