package endpoints

import (
	"context"
	"fmt"
	"net/url"
)

type (
	Competitor struct {
		Wins          int    `json:"wins"`
		Draws         int    `json:"draws"`
		Losses        int    `json:"losses"`
		Points        int    `json:"points"`
		Matches       int    `json:"matches"`
		TeamID        int    `json:"team_id"`
		Position      int    `json:"position"`
		TeamName      string `json:"team_name"`
		ScoresFor     int    `json:"scores_for"`
		ScoresAgainst int    `json:"scores_against"`
		TeamHashImage string `json:"team_hash_image"`
	}

	Standing struct {
		ID              int          `json:"id"`
		TournamentID    int          `json:"tournament_id"`
		Type            string       `json:"type"`
		Name            string       `json:"name"`
		SeasonID        int          `json:"season_id"`
		SeasonName      string       `json:"season_name"`
		LeagueID        int          `json:"league_id"`
		LeagueName      string       `json:"league_name"`
		LeagueHashImage string       `json:"league_hash_image"`
		Competitors     []Competitor `json:"competitors"`
	}
)

const Standings = "standings"

func (client *RestClient) GetStandings(ctx context.Context, sport string) ([]Standing, error) {
	return getStandings(client, ctx, sport, "", nil)
}

func (client *RestClient) GetStandingsByID(ctx context.Context, sport string, id int) ([]Standing, error) {
	return getStandings(client, ctx, sport, "id", id)
}

func (client *RestClient) GetStandingsByLeagueSeasonTypeID(ctx context.Context, sport string, id int, seasonID int, standingType string) ([]Standing, error) {
	standings := []Standing{}

	values := url.Values{}

	if id != 0 {
		values.Set("league_id", fmt.Sprintf("eq.%v", id))
	}

	if seasonID != 0 {
		values.Set("season_id", fmt.Sprintf("eq.%v", seasonID))
	}

	if standingType != "" {
		values.Set("type", fmt.Sprintf("eq.%v", standingType))
	}

	err := Get(client, ctx, sport, Standings, values, &standings)

	return standings, err
}

func getStandings(client *RestClient, ctx context.Context, sport string, query string, id any) ([]Standing, error) {
	standings := []Standing{}

	values := url.Values{}
	if query != "" && id != nil {
		values.Set(query, fmt.Sprintf("eq.%v", id))
	}

	err := Get(client, ctx, sport, Standings, values, &standings)

	return standings, err
}
