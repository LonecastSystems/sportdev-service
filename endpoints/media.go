package endpoints

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type Media struct {
	ID            int       `json:"id"`
	TeamID        int       `json:"team_id,omitempty"`
	PlayerID      int       `json:"player_id,omitempty"`
	LeagueID      int       `json:"league_id,omitempty"`
	Title         string    `json:"title"`
	Subtitle      string    `json:"subtitle"`
	URL           string    `json:"url"`
	ThumbnailURL  string    `json:"thumbnail_url"`
	DatePublished time.Time `json:"date_published"`
	ChannelURL    string    `json:"channel_url"`
}

const (
	MediaTeams   = "media-teams"
	MediaPlayers = "media-players"
	MediaLeagues = "media-leagues"
)

func (client *RestClient) GetMediaByTeamID(ctx context.Context, sport string, id int) ([]Media, error) {
	return getMedia(client, ctx, sport, MediaTeams, "team_id", id)
}

func (client *RestClient) GetMediaByPlayerID(ctx context.Context, sport string, id int) ([]Media, error) {
	return getMedia(client, ctx, sport, MediaPlayers, "player_id", id)
}

func (client *RestClient) GetMediaByLeagueID(ctx context.Context, sport string, id int) ([]Media, error) {
	return getMedia(client, ctx, sport, MediaLeagues, "league_id", id)
}

func getMedia(client *RestClient, ctx context.Context, sport string, resource string, query string, id int) ([]Media, error) {
	media := []Media{}

	values := url.Values{}
	values.Set(query, fmt.Sprintf("eq.%v", id))

	err := Get(client, ctx, sport, resource, values, &media)

	return media, err
}
