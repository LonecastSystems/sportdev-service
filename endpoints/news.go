package endpoints

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type (
	Subtitle struct {
		Text     string `json:"text"`
		Subtitle string `json:"subtitle"`
	}

	News struct {
		MatchID   int        `json:"match_id"`
		Date      time.Time  `json:"date"`
		Title     string     `json:"title"`
		Subtitles []Subtitle `json:"subtitles"`
	}
)

type AggNews struct {
	ID            int       `json:"id"`
	TeamID        int       `json:"team_id,omitempty"`
	MatchID       int       `json:"match_id,omitempty"`
	PlayerID      int       `json:"player_id,omitempty"`
	LeagueID      int       `json:"league_id,omitempty"`
	Title         string    `json:"title"`
	Link          string    `json:"link"`
	ThumbnailURL  string    `json:"thumbnail_url"`
	Description   string    `json:"description"`
	PublishedDate time.Time `json:"published_date"`
	SourceURL     string    `json:"source_url"`
	Source        string    `json:"source"`
}

const (
	NewsMatches    = "news-matches"
	AggNewsTeams   = "news-agg-teams"
	AggNewsMatches = "news-agg-matches"
	AggNewsLeagues = "news-agg-leagues"
	AggNewsPlayers = "news-agg-players"
)

func (client *RestClient) GetNewsByMatchID(ctx context.Context, sport string, id int) ([]News, error) {
	news := []News{}

	values := url.Values{}
	values.Set("match_id", fmt.Sprintf("eq.%v", id))

	err := Get(client, ctx, sport, NewsMatches, values, &news)

	return news, err
}

func (client *RestClient) GetAggNewsByTeamID(ctx context.Context, sport string, id int) ([]AggNews, error) {
	return getAggNews(client, ctx, sport, AggNewsTeams, "team_id", id)
}

func (client *RestClient) GetAggNewsByMatchID(ctx context.Context, sport string, id int) ([]AggNews, error) {
	return getAggNews(client, ctx, sport, AggNewsMatches, "match_id", id)
}

func (client *RestClient) GetAggNewsByLeagueID(ctx context.Context, sport string, id int) ([]AggNews, error) {
	return getAggNews(client, ctx, sport, AggNewsLeagues, "league_id", id)
}

func (client *RestClient) GetAggNewsByPlayerID(ctx context.Context, sport string, id int) ([]AggNews, error) {
	return getAggNews(client, ctx, sport, AggNewsPlayers, "player_id", id)
}

func getAggNews(client *RestClient, ctx context.Context, sport string, resource string, query string, id int) ([]AggNews, error) {
	news := []AggNews{}

	values := url.Values{}
	values.Set(query, fmt.Sprintf("eq.%v", id))

	err := Get(client, ctx, sport, resource, values, &news)

	return news, err
}
