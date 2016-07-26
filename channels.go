package twitch

// ChannelType -
type ChannelType struct {
	Mature                       bool              `json:"mature"`
	Status                       string            `json:"status"`
	BroadcasterLanguage          string            `json:"broadcaster_language"`
	DisplayName                  string            `json:"display_name"`
	Game                         string            `json:"game"`
	Delay                        int               `json:"delay"`
	Language                     string            `json:"language"`
	ID                           int               `json:"_id"`
	Name                         string            `json:"name"`
	CreatedAt                    string            `json:"created_at"`
	UpdatedAt                    string            `json:"updated_at"`
	Logo                         string            `json:"logo"`
	Banner                       string            `json:"banner"`
	VideoBanner                  string            `json:"video_banner"`
	Background                   string            `json:"background"`
	ProfileBanner                string            `json:"profile_banner"`
	ProfileBannerBackgroundColor string            `json:"profile_banner_background_color"`
	Partner                      bool              `json:"partner"`
	URL                          string            `json:"url"`
	Views                        int               `json:"views"`
	Followers                    int               `json:"followers"`
	Links                        map[string]string `json:"_links"`
}

//
// Implementation and their respective request/response types
//

// GetChannelInputType -
type GetChannelInputType struct {
	Channel string
}

// GetChannelOutputType -
type GetChannelOutputType ChannelType

// GetChannel -
func (session *Session) GetChannel(getChannelInputType *GetChannelInputType) (*GetChannelOutputType, error) {
	var out GetChannelOutputType
	err := session.request("GET", "/channels/"+getChannelInputType.Channel, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetChannelTeamsInputType -
type GetChannelTeamsInputType struct {
	Channel string
}

// GetChannelTeamsOutputType -
type GetChannelTeamsOutputType struct {
	Teams []TeamType        `json:"teams"`
	Links map[string]string `json:"_links"`
}

// GetChannelTeams -
func (session *Session) GetChannelTeams(getChannelTeamsInputType *GetChannelTeamsInputType) (*GetChannelTeamsOutputType, error) {
	var out GetChannelTeamsOutputType
	err := session.request("GET", "/channels/"+getChannelTeamsInputType.Channel+"/teams", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
