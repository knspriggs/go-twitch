package twitch

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

//
// Generic teams types
//

// TeamType -
type TeamType struct {
	ID          int               `json:"_id"`
	Name        string            `json:"name"`
	Info        string            `json:"info"`
	DisplayName string            `json:"display_name"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	Logo        string            `json:"logo"`
	Banner      string            `json:"banner"`
	Background  string            `json:"background"`
	Links       map[string]string `json:"_links"`
}

//
// Implementation and their respective request/response types
//

// GetAllTeamsInputType -
type GetAllTeamsInputType struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// GetAllTeamsOutputType -
type GetAllTeamsOutputType struct {
	Teams []TeamType        `json:"teams"`
	Links map[string]string `json:"_links"`
}

// GetAllTeams -
func (session *Session) GetAllTeams() (*GetAllTeamsOutputType, error) {
	u, err := url.Parse(session.URL + "/teams/")
	if err != nil {
		return &GetAllTeamsOutputType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GetAllTeamsOutputType{}, err
	}
	var sbc GetAllTeamsOutputType
	err = json.Unmarshal([]byte(body), &sbc)
	if err != nil {
		return &GetAllTeamsOutputType{}, err
	}

	return &sbc, nil
}

// GetTeamInputType -
type GetTeamInputType struct {
	Team string
}

// GetTeamOutputType -
type GetTeamOutputType TeamType

// GetTeam -
func (session *Session) GetTeam(getTeamInputType *GetTeamInputType) (*GetTeamOutputType, error) {
	u, err := url.Parse(session.URL + "/teams/" + getTeamInputType.Team)
	if err != nil {
		return &GetTeamOutputType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GetTeamOutputType{}, err
	}
	var sbc GetTeamOutputType
	err = json.Unmarshal([]byte(body), &sbc)
	if err != nil {
		return &GetTeamOutputType{}, err
	}

	return &sbc, nil
}
