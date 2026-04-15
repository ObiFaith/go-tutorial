package services

import (
	"encoding/json"
	"genderize-api/config"
	"io"
	"net/http"
	"net/url"
	"time"
)

type GenderResponse struct {
	Count int	`json:"count"`
	Name string	`json:"name"`
	Gender string	`json:"gender"`
	Probability float64	`json:"probability"`
}

func FetchGenderResponse(name string) (GenderResponse, error) {
	cfg := config.LoadConfig()
	genderizeApi, _ := url.Parse(cfg.GenderizeApi)

	params := url.Values{}
	params.Add("name", name)

	genderizeApi.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", genderizeApi.String(), nil)

	if err != nil {
		return GenderResponse{}, err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return GenderResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return GenderResponse{}, err
	}

	var data GenderResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return GenderResponse{}, err
	}

	return data, nil
}