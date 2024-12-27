package tests

import (
	"context"
	"testing"
)

func TestGetInjuriesByPlayerID(t *testing.T) {
	c := CreateClient(t)

	id := 21880

	injuries, err := c.GetInjuriesByPlayerID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(injuries); len == 0 {
		t.Fatal(len)
	}

	if injuries[0].PlayerID != id {
		t.Fail()
	}
}

func TestGetInjuriesByMatchID(t *testing.T) {
	c := CreateClient(t)

	id := 28911

	injuries, err := c.GetInjuriesByMatchID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(injuries); len == 0 {
		t.Fatal(len)
	}

	if injuries[0].MatchID != id {
		t.Fail()
	}
}

func TestGetInjuriesBySeasonID(t *testing.T) {
	c := CreateClient(t)

	id := 7172

	injuries, err := c.GetInjuriesBySeasonID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(injuries); len == 0 {
		t.Fatal(len)
	}

	if injuries[0].SeasonID != id {
		t.Fail()
	}
}

func TestGetInjuriesByTournamentID(t *testing.T) {
	c := CreateClient(t)

	id := 269

	injuries, err := c.GetInjuriesByTournamentID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(injuries); len == 0 {
		t.Fatal(len)
	}

	if injuries[0].TournamentID != id {
		t.Fail()
	}
}
