package twitch

// Game - describes a game on twitch
type Game struct {
	GameInfo GameInfo `json:"game"`
	Viewers  int      `json:"viewers"`
	Channels int      `json:"channels"`
}

// GameInfo - details about the specific game
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

// GetTopGamesInputType - request type for GetTopGames
type GetTopGamesInputType struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// GetTopGamesOutputType - response type containing an array of games
type GetTopGamesOutputType struct {
	Links map[string]string `json:"_links"`
	Total int               `json:"_total"`
	Top   []Game            `json:"top"`
}

// GetTopGames - returns the top games at the time of request on twitch
func (session *Session) GetTopGames(getTopeGamesInputType *GetTopGamesInputType) (*GetTopGamesOutputType, error) {
	var out GetTopGamesOutputType
	err := session.request("GET", "/games/top", &getTopeGamesInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
