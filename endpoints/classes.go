package endpoints

import (
	"context"
	"fmt"
	"net/url"
)

type Class struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Flag       string `json:"flag"`
	Alpha      string `json:"alpha"`
	Importance int    `json:"importance"`
	HashImage  string `json:"hash_image"`
}

const Classes = "classes"

func (client *RestClient) GetClasses(ctx context.Context, sport string) ([]Class, error) {
	classes := []Class{}

	err := Get(client, ctx, sport, Classes, url.Values{}, &classes)

	return classes, err
}

func (client *RestClient) GetClassesByID(ctx context.Context, sport string, id int) ([]Class, error) {
	classes := []Class{}

	values := url.Values{}
	values.Set("id", fmt.Sprintf("eq.%v", id))

	err := Get(client, ctx, sport, Classes, values, &classes)

	return classes, err
}

func (client *RestClient) GetClassesByAlpha(ctx context.Context, sport string, alpha string) ([]Class, error) {
	classes := []Class{}

	values := url.Values{}
	values.Set("alpha", fmt.Sprintf("eq.%v", alpha))

	err := Get(client, ctx, sport, Classes, values, &classes)

	return classes, err
}
