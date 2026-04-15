package services

import (
	"encoding/json"
	"genderize-api/config"
	"io"
	"net/http"
	"net/url"
	"time"
)

type AgeResponse struct {
	Age int `json:"age"`
}

func FetchAgeResponse(name string) (AgeResponse, error) {
	cfg := config.LoadConfig()
	agifyApi, _ := url.Parse(cfg.AgifyApi)

	params := url.Values{}
	params.Add("name", name)

	agifyApi.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", agifyApi.String(), nil)

	if err != nil {
		return AgeResponse{}, err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return AgeResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AgeResponse{}, err
	}

	var data AgeResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return AgeResponse{}, err
	}

	return data, nil
}