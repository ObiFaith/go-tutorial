package clients

import (
	"context"
	"net/url"
)

type GenderResponse struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type AgeResponse struct {
	Age int `json:"age"`
}

type Country struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type NationalizeResponse struct {
	Country []Country `json:"country"`
}

func FetchGender(ctx context.Context, baseURL, name string) (GenderResponse, error) {
	return FetchJSON[GenderResponse](ctx, baseURL, url.Values{"name": {name}})
}

func FetchAge(ctx context.Context, baseURL, name string) (AgeResponse, error) {
	return FetchJSON[AgeResponse](ctx, baseURL, url.Values{"name": {name}})
}

