package twitch

//
// Generic teams types
//

// TeamType - describes a team on twitch
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

// GetAllTeamsInputType - request parameters for the GetAllTeams function
type GetAllTeamsInputType struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// GetAllTeamsOutputType - contains a list of teams
type GetAllTeamsOutputType struct {
	Teams []TeamType        `json:"teams"`
	Links map[string]string `json:"_links"`
}

// GetAllTeams - returns all the teams on twitch
func (session *Session) GetAllTeams() (*GetAllTeamsOutputType, error) {
	var out GetAllTeamsOutputType
	err := session.request("GET", "/teams", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetTeamInputType - specifies the team name
type GetTeamInputType struct {
	Team string
}

// GetTeamOutputType - containers the team
type GetTeamOutputType TeamType

// GetTeam - returns the specified team
func (session *Session) GetTeam(getTeamInputType *GetTeamInputType) (*GetTeamOutputType, error) {
	var out GetTeamOutputType
	err := session.request("GET", "/teams/"+getTeamInputType.Team, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
