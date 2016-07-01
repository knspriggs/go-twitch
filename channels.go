package twitch

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

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

type GetChannelInputType struct {
	Channel string
}
type GetChannelOutputType ChannelType

func (session *Session) GetChannel(getChannelInputType *GetChannelInputType) (*GetChannelOutputType, error) {
	u, err := url.Parse(session.URL + "/channels/" + getChannelInputType.Channel)
	if err != nil {
		return &GetChannelOutputType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GetChannelOutputType{}, err
	}
	var out GetChannelOutputType
	err = json.Unmarshal([]byte(body), &out)
	if err != nil {
		return &GetChannelOutputType{}, err
	}

	return &out, nil
}
