package services

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Data struct {
	Count int
	Name string
	Gender string
	Probability float64
}

func FetchGenderData(name string) (Data, error) {
	req, err := http.NewRequest("GET", "https://api.genderize.io/?name="+name, nil)
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