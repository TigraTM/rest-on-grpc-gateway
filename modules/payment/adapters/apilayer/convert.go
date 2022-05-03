package apilayer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"rest-on-grpc-gateway/modules/payment/internal/app"

	"github.com/shopspring/decimal"
)

type convertResponse struct {
	Result decimal.Decimal `json:"result"`
}

const convertURL = "/convert"

// ConvertAmount converts the amount  from currency to currency.
func (c *Client) ConvertAmount(ctx context.Context, fromCurrency, toCurrency string, amount decimal.Decimal) (decimal.Decimal, error) {
	uri, err := c.getConvertURL(fromCurrency, toCurrency, amount)
	if err != nil {
		return decimal.Decimal{}, fmt.Errorf("c.getConvertUrl: %w", err)
	}

	respBody, err := c.request(ctx, uri, http.MethodGet, nil)
	switch {
	case errors.Is(err, errUnknown), errors.Is(err, errNotSuccess):
		return decimal.Decimal{}, app.ErrExchangeClient
	case err != nil:
		return decimal.Decimal{}, fmt.Errorf("c.request: %w", err)
	}

	result := &convertResponse{}
	if err = json.Unmarshal(respBody, result); err != nil {
		return decimal.Decimal{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result.Result, nil
}

func (c *Client) getConvertURL(fromCurrency, toCurrency string, amount decimal.Decimal) (string, error) {
	str := fmt.Sprintf("%s%s", c.basePath, convertURL)

	uri, err := url.Parse(str)
	if err != nil {
		return "", fmt.Errorf("url.Parse: %w", err)
	}

	parameters := url.Values{}

	parameters.Add("apikey", c.apiKey)
	parameters.Add("from", fromCurrency)
	parameters.Add("to", toCurrency)
	parameters.Add("amount", amount.String())

	uri.RawQuery = parameters.Encode()

	return uri.String(), nil
}
