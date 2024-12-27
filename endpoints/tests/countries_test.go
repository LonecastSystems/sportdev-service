package tests

import (
	"context"
	"testing"
)

func TestGetCountries(t *testing.T) {
	c := CreateClient(t)

	countries, err := c.GetCountries(context.Background(), "football")

	if err != nil {
		t.Fatal(err)
	}

	if len := len(countries); len == 0 {
		t.Fatal(len)
	}
}

func TestGetCountriesByID(t *testing.T) {
	c := CreateClient(t)

	id := 1

	countries, err := c.GetCountriesByID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(countries); len == 0 {
		t.Fatal(len)
	}

	if countries[0].ID != id {
		t.Fail()
	}
}

func TestGetCountriesByAlpha(t *testing.T) {
	c := CreateClient(t)

	alpha := "CA"

	countries, err := c.GetCountriesByAlpha(context.Background(), "football", alpha)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(countries); len == 0 {
		t.Fatal(len)
	}

	if countries[0].Alpha != alpha {
		t.Fail()
	}
}
