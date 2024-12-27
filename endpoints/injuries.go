package endpoints

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type Injury struct {
	PlayerID        int       `json:"player_id"`
	PlayerName      string    `json:"player_name"`
	PlayerHashImage string    `json:"player_hash_image"`
	Type            string    `json:"type"`
	Reason          string    `json:"reason"`
	MatchID         int       `json:"match_id"`
	SeasonID        int       `json:"season_id"`
	TournamentID    int       `json:"tournament_id"`
	StartTimestamp  time.Time `json:"start_timestamp"`
}

const Injuries = "injuries"

func (client *RestClient) GetInjuriesByPlayerID(ctx context.Context, sport string, id int) ([]Injury, error) {
	return getInjuries(client, ctx, sport, "player_id", id)
}

func (client *RestClient) GetInjuriesByMatchID(ctx context.Context, sport string, id int) ([]Injury, error) {
	return getInjuries(client, ctx, sport, "match_id", id)
}

func (client *RestClient) GetInjuriesBySeasonID(ctx context.Context, sport string, id int) ([]Injury, error) {
	return getInjuries(client, ctx, sport, "season_id", id)
}

func (client *RestClient) GetInjuriesByTournamentID(ctx context.Context, sport string, id int) ([]Injury, error) {
	return getInjuries(client, ctx, sport, "tournament_id", id)
}

func getInjuries(client *RestClient, ctx context.Context, sport string, query string, id int) ([]Injury, error) {
	injuries := []Injury{}

	values := url.Values{}
	values.Set(query, fmt.Sprintf("eq.%v", id))

	err := Get(client, ctx, sport, Injuries, values, &injuries)

	return injuries, err
}
