package services

import (
	"encoding/json"
	"fmt"
	"genderize-api/config"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Country struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type NationalizeResponse struct {
	Count   int       `json:"count"`
	Name    string    `json:"name"`
	Country []Country `json:"country"`
}

func FetchCountryResponse(name string) (Country, error) {
	cfg := config.LoadConfig()
	nationalizeApi, _ := url.Parse(cfg.NationalizeApi)

	params := url.Values{}
	params.Add("name", name)
	nationalizeApi.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", nationalizeApi.String(), nil)
	if err != nil {
		return Country{}, err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return Country{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Country{}, err
	}

	var data NationalizeResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return Country{}, err
	}

	if len(data.Country) == 0 {
		return Country{}, fmt.Errorf("no country data found")
	}

	highest := data.Country[0]
	for _, c := range data.Country {
		if c.Probability > highest.Probability {
			highest = c
		}
	}

	return highest, nil
}
