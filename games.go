package twitch

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/google/go-querystring/query"
)

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

//
// Implementation and their respective request/response types
//

type GetTopGamesInputType struct {
	Limit  int
	Offset int
}

type GetTopGamesOutputType struct {
	Links map[string]string `json:"_links"`
	Total int               `json:"_total"`
	Top   []Game            `json:"top"`
}

func (session *Session) GetTopGames(getTopeGamesInputType *GetTopGamesInputType) (*GetTopGamesOutputType, error) {
	q, err := query.Values(getTopeGamesInputType)
	if err != nil {
		return &GetTopGamesOutputType{}, err
	}
	u, err := url.Parse(session.URL + "/games/top?" + q.Encode())
	if err != nil {
		return &GetTopGamesOutputType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GetTopGamesOutputType{}, err
	}
	var grt GetTopGamesOutputType
	err = json.Unmarshal([]byte(body), &grt)
	if err != nil {
		return &GetTopGamesOutputType{}, err
	}

	return &grt, nil
}
