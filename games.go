package twitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type GamesRequestType struct {
	Limit  int
	Offset int
}

type GamesResponseType struct {
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

func (session *Session) GetTopGames(gamesRequestType *GamesRequestType) (*GamesResponseType, error) {
	u, err := url.Parse(session.URL + "/games/top")
	if err != nil {
		return &GamesResponseType{}, err
	}
	q := u.Query()
	q.Set("limit", fmt.Sprintf("%d", gamesRequestType.Limit))
	q.Set("offset", fmt.Sprintf("%d", gamesRequestType.Offset))
	u.RawQuery = q.Encode()

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GamesResponseType{}, err
	}
	var grt GamesResponseType
	err = json.Unmarshal([]byte(body), &grt)
	if err != nil {
		return &GamesResponseType{}, err
	}

	return &grt, nil
}
