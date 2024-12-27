package tests

import (
	"context"
	"testing"
)

func TestGetStandings(t *testing.T) {
	c := CreateClient(t)

	standings, err := c.GetStandings(context.Background(), "football")

	if err != nil {
		t.Fatal(err)
	}

	if len := len(standings); len == 0 {
		t.Fatal(len)
	}
}

func TestGetStandingsByID(t *testing.T) {
	c := CreateClient(t)

	id := 1

	standings, err := c.GetStandingsByID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(standings); len == 0 {
		t.Fatal(len)
	}

	if standings[0].ID != id {
		t.Fail()
	}
}

func TestGetStandingsByLeagueSeasonTypeID(t *testing.T) {
	c := CreateClient(t)

	id := 29
	seasonID := 2
	standingType := "total"

	standings, err := c.GetStandingsByLeagueSeasonTypeID(context.Background(), "football", id, seasonID, standingType)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(standings); len == 0 {
		t.Fatal(len)
	}

	if standings[0].LeagueID != id {
		t.Fail()
	}

	if standings[0].SeasonID != seasonID {
		t.Fail()
	}

	if standings[0].Type != standingType {
		t.Fail()
	}
}
