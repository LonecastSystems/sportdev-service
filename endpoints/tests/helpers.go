package tests

import (
	"testing"

	"github.com/LonecastSystems/sportdev-go/endpoints"
)

const apiKey = ""

func CreateClient(t *testing.T) *endpoints.RestClient {
	if apiKey == "" {
		t.Skip("Invalid credentials")
	}

	return endpoints.CreateClient(apiKey)
}
