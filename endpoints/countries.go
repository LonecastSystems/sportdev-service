package endpoints

import (
	"context"
	"fmt"
	"net/url"
)

type Country struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Alpha     string `json:"alpha"`
	HashImage string `json:"hash_image"`
}

const Countries = "countries"

func (client *RestClient) GetCountries(ctx context.Context, sport string) ([]Country, error) {
	countries := []Country{}

	err := Get(client, ctx, sport, Countries, url.Values{}, &countries)

	return countries, err
}

func (client *RestClient) GetCountriesByID(ctx context.Context, sport string, id int) ([]Country, error) {
	countries := []Country{}

	values := url.Values{}
	values.Set("id", fmt.Sprintf("eq.%v", id))

	err := Get(client, ctx, sport, Countries, values, &countries)

	return countries, err
}

func (client *RestClient) GetCountriesByAlpha(ctx context.Context, sport string, alpha string) ([]Country, error) {
	countries := []Country{}

	values := url.Values{}
	values.Set("alpha", fmt.Sprintf("eq.%v", alpha))

	err := Get(client, ctx, sport, Countries, values, &countries)

	return countries, err
}
