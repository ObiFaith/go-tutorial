package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var HTTPClient = &http.Client{
	Timeout: 10 * time.Second,
}

func FetchJSON[T any](ctx context.Context, baseUrl string, params url.Values) (T, error) {
	var result T

	parseUrl, err := url.Parse(baseUrl)
	if err != nil {
		return result, fmt.Errorf("invalid url: %w", err)
	}

	parseUrl.RawQuery = params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parseUrl.String(), nil)

	if err != nil {
		return result, err
	}

	res, err := HTTPClient.Do(req)

	if err != nil {
		return result, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return result, fmt.Errorf("external api failed with status: %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}