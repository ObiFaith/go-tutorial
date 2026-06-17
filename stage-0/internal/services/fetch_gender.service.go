package services

import (
	"context"
	"encoding/json"
	"genderize-api/config"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Data struct {
	Count       int
	Name        string
	Gender      string
	Probability float64
}

var (
	httpClient *http.Client
	apiBaseURL *url.URL
	once       sync.Once
)

func initClient() {
	once.Do(func() {
		httpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
		cfg := config.LoadConfig()
		apiBaseURL, _ = url.Parse(cfg.GenderizeApi)
	})
}

func FetchGenderData(ctx context.Context, name string) (Data, error) {
	initClient()

	u := *apiBaseURL
	params := url.Values{}
	params.Add("name", name)

	u.RawQuery = params.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)

	if err != nil {
		return Data{}, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return Data{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Data{}, err
	}

	var data Data
	if err := json.Unmarshal(body, &data); err != nil {
		return Data{}, err
	}

	return data, nil
}
