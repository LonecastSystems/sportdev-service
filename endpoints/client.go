package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type RestClient struct {
	Client *http.Client
	ApiKey string
}

const (
	Offset = "offset"
	Limit  = "limit"
	Lang   = "lang"
)

func CreateClient(ApiKey string) *RestClient {
	return &RestClient{Client: &http.Client{}, ApiKey: ApiKey}
}

func (client *RestClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", client.ApiKey))

	return client.Client.Do(req)
}

func Get[T any](client *RestClient, ctx context.Context, sport string, resource string, values url.Values, response *T) error {
	var offset string
	if cOffset := ctx.Value(Offset); cOffset != nil {
		offset = cOffset.(string)
	} else {
		offset = "0"
	}

	var limit string
	if cLimit := ctx.Value(Limit); cLimit != nil {
		limit = cLimit.(string)
	} else {
		limit = "50"
	}

	var lang string
	if cLang := ctx.Value(Lang); cLang != nil {
		lang = cLang.(string)
	} else {
		lang = "en"
	}

	values.Set("offset", offset)
	values.Set("limit", limit)
	values.Set("lang", lang)

	path := fmt.Sprintf("https://%v.sportdevs.com/%v", sport, resource)

	url := url.URL{Path: path, RawQuery: values.Encode()}
	urlStr := url.RequestURI()

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%v failed: %v", urlStr, res.StatusCode)
	}

	defer res.Body.Close()
	if body, err := io.ReadAll(res.Body); err != nil {
		return err
	} else {
		return json.Unmarshal(body, &response)
	}
}
