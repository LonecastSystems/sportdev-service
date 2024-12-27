package tests

import (
	"context"
	"testing"
)

func TestGetNewsByMatchID(t *testing.T) {
	c := CreateClient(t)

	id := 26100

	news, err := c.GetNewsByMatchID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(news); len == 0 {
		t.Fatal(len)
	}

	if news[0].MatchID != id {
		t.Fail()
	}
}

func TestGetAggNewsByTeamID(t *testing.T) {
	c := CreateClient(t)

	id := 19829

	news, err := c.GetAggNewsByTeamID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(news); len == 0 {
		t.Fatal(len)
	}

	if news[0].TeamID != id {
		t.Fail()
	}
}

func TestGetAggNewsByMatchID(t *testing.T) {
	c := CreateClient(t)

	id := 1911869

	news, err := c.GetAggNewsByMatchID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(news); len == 0 {
		t.Fatal(len)
	}

	if news[0].TeamID != id {
		t.Fail()
	}
}

func TestGetAggNewsByLeagueID(t *testing.T) {
	c := CreateClient(t)

	id := 9963

	news, err := c.GetAggNewsByLeagueID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(news); len == 0 {
		t.Fatal(len)
	}

	if news[0].LeagueID != id {
		t.Fail()
	}
}

func TestGetAggNewsByPlayerID(t *testing.T) {
	c := CreateClient(t)

	id := 213836

	news, err := c.GetAggNewsByPlayerID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(news); len == 0 {
		t.Fatal(len)
	}

	if news[0].PlayerID != id {
		t.Fail()
	}
}
