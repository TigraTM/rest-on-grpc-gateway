package apilayer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"rest-on-grpc-gateway/modules/payment/internal/app"
)

type symbols struct {
	Symbols map[string]string `json:"symbols"`
}

const symbolsURL = "/symbols"

// GetSymbols get all symbols.
func (c *Client) GetSymbols(ctx context.Context) (map[string]string, error) {
	uri, err := c.getSymbolsURL()
	if err != nil {
		return nil, fmt.Errorf("c.getSymbolsUrl: %w", err)
	}

	respBody, err := c.request(ctx, uri, http.MethodGet, nil)
	switch {
	case errors.Is(err, errUnknown), errors.Is(err, errNotSuccess):
		return nil, app.ErrExchangeClient
	case err != nil:
		return nil, fmt.Errorf("c.request: %w", err)
	}

	result := &symbols{}
	if err = json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result.Symbols, nil
}

func (c *Client) getSymbolsURL() (string, error) {
	str := fmt.Sprintf("%s%s", c.basePath, symbolsURL)

	uri, err := url.Parse(str)
	if err != nil {
		return "", fmt.Errorf("url.Parse: %w", err)
	}

	parameters := url.Values{}

	parameters.Add("apikey", c.apiKey)

	uri.RawQuery = parameters.Encode()

	return uri.String(), nil
}
