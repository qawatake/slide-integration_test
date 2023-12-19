package greeting

import (
	"context"
	"errors"
	"net/http"
	"sample_app/config"
	"strings"
)

type Client struct {
	cfg        config.GreetingConfig
	httpClient *http.Client
}

func New(httpClient *http.Client, cfg config.GreetingConfig) *Client {
	return &Client{
		httpClient: httpClient,
		cfg:        cfg,
	}
}

func NewHTTPClient() *http.Client {
	return &http.Client{}
}

func (c *Client) Hello(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.cfg.URL+"/hello", strings.NewReader(`Hello, World!`))
	if err != nil {
		return err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	switch res.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return errors.New("unexpected status code")
	}
}
