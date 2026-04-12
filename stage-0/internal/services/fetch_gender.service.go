package services

import (
	"encoding/json"
	"genderize-api/config"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Data struct {
	Count int
	Name string
	Gender string
	Probability float64
}

func FetchGenderData(name string) (Data, error) {
	cfg := config.LoadConfig()
	genderizeApi, _ := url.Parse(cfg.GenderizeApi)

	params := url.Values{}
	params.Add("name", name)

	genderizeApi.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", genderizeApi.String(), nil)

	if err != nil {
		return Data{}, err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
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