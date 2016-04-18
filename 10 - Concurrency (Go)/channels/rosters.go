package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var apiToken string

type stattleshipTeamsResponse struct {
	// Has other fields but we only care about "teams"
	Teams []team `json:"teams"`
}

type stattleshipRosterResponse struct {
	Players []player `json:"players"`
}

type team struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type player struct {
	Name string `json:"name"`
}

func (p *player) String() string {
	return p.Name
}

func init() {
	apiToken = os.Getenv("STATTLESHIP_API_TOKEN")
}

func newRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", apiToken)
	req.Header.Set("Accept", "application/vnd.stattleship.com; version=1")
	return req, nil
}

func do(req *http.Request) (io.ReadCloser, error) {
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status: %s", res.Status)
	}

	return res.Body, nil
}

func getTeams() ([]team, error) {
	req, err := newRequest("https://www.stattleship.com/football/nfl/teams")

	if err != nil {
		return nil, err
	}

	body, err := do(req)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	var m stattleshipTeamsResponse
	if err := json.NewDecoder(body).Decode(&m); err != nil {
		return nil, err
	}
	return m.Teams, nil
}

func getRoster(slug string) ([]player, error) {
	url := fmt.Sprintf("https://www.stattleship.com/football/nfl/rosters?team_id=%s", slug)
	req, err := newRequest(url)
	if err != nil {
		return nil, err
	}

	body, err := do(req)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	var m stattleshipRosterResponse
	if err := json.NewDecoder(body).Decode(&m); err != nil {
		return nil, err
	}
	return m.Players, nil
}

func main() {
	// Get list of teams
	teams, err := getTeams()
	if err != nil {
		log.Fatal(err)
	}

	// For each team, get roster and print
	for _, t := range teams {
		players, err := getRoster(t.Slug)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(t.Name)
		for _, p := range players {
			fmt.Print(p)
		}
		fmt.Print("\n\n")
	}
}
