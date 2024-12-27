package tests

import (
	"context"
	"testing"
)

func TestGetMediaByTeamID(t *testing.T) {
	c := CreateClient(t)

	id := 14

	media, err := c.GetMediaByTeamID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(media); len == 0 {
		t.Fatal(len)
	}

	if media[0].TeamID != id {
		t.Fail()
	}
}

func TestGetMediaByPlayerID(t *testing.T) {
	c := CreateClient(t)

	id := 328523

	media, err := c.GetMediaByPlayerID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(media); len == 0 {
		t.Fatal(len)
	}

	if media[0].PlayerID != id {
		t.Fail()
	}
}

func TestGetMediaByLeagueID(t *testing.T) {
	c := CreateClient(t)

	id := 21

	media, err := c.GetMediaByLeagueID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(media); len == 0 {
		t.Fatal(len)
	}

	if media[0].LeagueID != id {
		t.Fail()
	}
}
