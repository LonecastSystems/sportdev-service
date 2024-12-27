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
	return getClasses(client, ctx, sport, "", nil)
}

func (client *RestClient) GetClassesByID(ctx context.Context, sport string, id int) ([]Class, error) {
	return getClasses(client, ctx, sport, "id", id)
}

func (client *RestClient) GetClassesByAlpha(ctx context.Context, sport string, alpha string) ([]Class, error) {
	return getClasses(client, ctx, sport, "alpha", alpha)
}

func getClasses(client *RestClient, ctx context.Context, sport string, query string, id any) ([]Class, error) {
	classes := []Class{}

	values := url.Values{}
	if query != "" && id != nil {
		values.Set(query, fmt.Sprintf("eq.%v", id))
	}

	err := Get(client, ctx, sport, Classes, values, &classes)

	return classes, err
}
