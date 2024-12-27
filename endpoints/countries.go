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
	return getCountries(client, ctx, sport, "", nil)
}

func (client *RestClient) GetCountriesByID(ctx context.Context, sport string, id int) ([]Country, error) {
	return getCountries(client, ctx, sport, "id", id)
}

func (client *RestClient) GetCountriesByAlpha(ctx context.Context, sport string, alpha string) ([]Country, error) {
	return getCountries(client, ctx, sport, "alpha", alpha)
}

func getCountries(client *RestClient, ctx context.Context, sport string, query string, id any) ([]Country, error) {
	countries := []Country{}

	values := url.Values{}
	if query != "" && id != nil {
		values.Set(query, fmt.Sprintf("eq.%v", id))
	}

	err := Get(client, ctx, sport, Countries, values, &countries)

	return countries, err
}
