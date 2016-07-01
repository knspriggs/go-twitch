package twitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type GamesInputType struct {
	Limit  int
	Offset int
}

type GamesOutputType struct {
	Links map[string]string `json:"_links"`
	Total int               `json:"_total"`
	Top   []Game            `json:"top"`
}

type Game struct {
	GameInfo GameInfo `json:"game"`
	Viewers  int      `json:"viewers"`
	Channels int      `json:"channels"`
}

type GameInfo struct {
	Name        string            `json:"name"`
	Box         map[string]string `json:"box"`
	Logo        map[string]string `json:"logo"`
	Links       map[string]string `json:"_links"`
	ID          int               `json:"_id"`
	GiantBombID int               `json:"giantbomb_id"`
}

func (session *Session) GetTopGames(gamesInputType *GamesInputType) (*GamesOutputType, error) {
	u, err := url.Parse(session.URL + "/games/top")
	if err != nil {
		return &GamesOutputType{}, err
	}
	q := u.Query()
	q.Set("limit", fmt.Sprintf("%d", gamesInputType.Limit))
	q.Set("offset", fmt.Sprintf("%d", gamesInputType.Offset))
	u.RawQuery = q.Encode()

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GamesOutputType{}, err
	}
	var grt GamesOutputType
	err = json.Unmarshal([]byte(body), &grt)
	if err != nil {
		return &GamesOutputType{}, err
	}

	return &grt, nil
}
