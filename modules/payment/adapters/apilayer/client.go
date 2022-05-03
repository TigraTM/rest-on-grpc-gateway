// Package apilayer contains a client to convert currencies.
package apilayer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	clientTimeout = time.Second * 10
)

// Errors ...
var (
	errUnknown    = errors.New("unknown")
	errNotSuccess = errors.New("not success")
)

// Client structure for interacting with APILayer, for convert currencies.
type Client struct {
	apiKey   string
	basePath string

	client *http.Client
}

// New build and return new APILayer client.
func New(apiKey, basePath string) *Client {
	client := &http.Client{
		Timeout: clientTimeout,
	}

	return &Client{
		apiKey:   apiKey,
		basePath: basePath,
		client:   client,
	}
}

func (c *Client) request(ctx context.Context, url, method string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client do: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	base := &BaseAPILayer{}
	if err = json.Unmarshal(respBody, base); err != nil {
		return nil, err
	}

	if !base.Success {
		return nil, fmt.Errorf("%w: reason - %s", errNotSuccess, base.Error.Info)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return respBody, nil
	default:
		return nil, errUnknown
	}
}
