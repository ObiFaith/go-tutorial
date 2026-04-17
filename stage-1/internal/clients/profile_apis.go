package clients

import (
	"context"
	"fmt"
	"net/url"
)

type Client struct {
	GenderizeUrl 		string
	AgifyUrl    		string
	NationalizeUrl	string
}

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

func (c *Client) FetchGender(ctx context.Context, name string) (GenderResponse, error) {
	return FetchJSON[GenderResponse](ctx, c.GenderizeUrl, url.Values{"name": {name}})
}

func (c *Client) FetchAge(ctx context.Context, name string) (AgeResponse, error) {
	return FetchJSON[AgeResponse](ctx, c.AgifyUrl, url.Values{"name": {name}})
}

func (c *Client) FetchCountry(ctx context.Context, name string) (Country, error) {
	res, err := FetchJSON[NationalizeResponse](ctx, c.NationalizeUrl, url.Values{"name": {name}})

	if err != nil {
		return Country{}, err
	}
	if (len(res.Country) == 0){
		return Country{}, fmt.Errorf("No country found!")
	}

	highest := res.Country[0]

	for _, country := range res.Country[1:]{
		if country.Probability > highest.Probability{
			highest = country
		}
	}

	return highest, nil
}

