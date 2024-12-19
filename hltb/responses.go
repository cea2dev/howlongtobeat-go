package hltb

import (
	"fmt"
)

type searchResponse struct {
	Games []struct {
		GameId         int64  `json:"game_id"`
		Name           string `json:"game_name"`
		Alias          string `json:"game_alias"`
		Image          string `json:"game_image"`
		Developer      string `json:"profile_dev"`
		Platforms      string `json:"profile_platform"`
		SteamProfileId int64
	} `json:"data"`
}

type GameEntry struct {
	Id             int64
	Name           string
	Alias          string
	ImageUrl       string
	Developer      string
	Platforms      []string
	SteamProfileId *int64
}

func (r *searchResponse) getGameEntries(baseImageUrl, term string) []GameEntry {
	res := []GameEntry{}

	for _, g := range r.Games {
		res = append(res, GameEntry{
			Id:             g.GameId,
			Name:           g.Name,
			Alias:          g.Alias,
			ImageUrl:       fmt.Sprintf("%s/%s", baseImageUrl, g.Image),
			Platforms:      splitStrTerms(g.Platforms, ","),
			SteamProfileId: getSteamProfileId(g.SteamProfileId),
		})
	}

	return res
}

func getSteamProfileId(id int64) *int64 {
	if id > 0 {
		return &id
	}

	return nil
}
