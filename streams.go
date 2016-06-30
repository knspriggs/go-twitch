package twitch

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/google/go-querystring/query"
)

//
// Generic streams types
//

type StreamType struct {
	Game        string            `json:"game"`
	Viewers     int               `json:"viewers"`
	AverageFPS  float32           `json:"average_fps"`
	Delay       int               `json:"delay"`
	VideoHeight int               `json:"video_height"`
	IsPlaylist  bool              `json:"is_playlist"`
	CreatedAt   string            `json:"created_at"`
	ID          int               `json:"_id"`
	Channel     ChannelType       `json:"channel"`
	Preview     map[string]string `json:"preview"`
	Links       map[string]string `json:"_links"`
}

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

type FeaturedType struct {
	Image     string     `json:"image"`
	Text      string     `json:"text"`
	Title     string     `json:"title"`
	Sponsored bool       `json:"sponsored"`
	Scheduled bool       `json:"scheduled"`
	Stream    StreamType `json:"stream"`
}

//
// Implementation and their respective request/response types
//

type GetStreamsRequestType struct {
	Game       string `url:"game,omitempty"`
	Channel    string `url:"channel,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
	ClientID   string `url:"client_id,omitempty"`
	StreamType string `url:"stream_type,omitempty"`
	Language   string `url:"language,omitempty"`
}

type GetStreamsResponseType struct {
	Total   int               `json:"_total"`
	Streams []StreamType      `json:"streams"`
	Links   map[string]string `json:"_links"`
}

func (session *Session) GetStream(getStreamsRequestType *GetStreamsRequestType) (*GetStreamsResponseType, error) {
	q, err := query.Values(getStreamsRequestType)
	if err != nil {
		return &GetStreamsResponseType{}, err
	}
	u, err := url.Parse(session.URL + "/streams?" + q.Encode())
	if err != nil {
		return &GetStreamsResponseType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GetStreamsResponseType{}, err
	}
	var s GetStreamsResponseType
	err = json.Unmarshal([]byte(body), &s)
	if err != nil {
		return &GetStreamsResponseType{}, err
	}

	return &s, nil
}

type GetStreamByChannelRequestType struct {
	Channel string `url:"channel"`
}

type GetStreamByChannelResponseType struct {
	Stream StreamType        `json:"stream"`
	Links  map[string]string `json:"_links"`
}

func (session *Session) GetStreamByChannel(getStreamByChannelRequestType *GetStreamByChannelRequestType) (*GetStreamByChannelResponseType, error) {
	u, err := url.Parse(session.URL + "/streams/" + getStreamByChannelRequestType.Channel)
	if err != nil {
		return &GetStreamByChannelResponseType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GetStreamByChannelResponseType{}, err
	}
	var sbc GetStreamByChannelResponseType
	err = json.Unmarshal([]byte(body), &sbc)
	if err != nil {
		return &GetStreamByChannelResponseType{}, err
	}

	return &sbc, nil
}

type GetFeaturedStreamsRequestType struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type GetFeaturedStreamsResponseType struct {
	Featured []FeaturedType    `json:"featured"`
	Links    map[string]string `json:"_links"`
}

func (session *Session) GetFeaturedStreams(getFeaturedStreamsRequestType *GetFeaturedStreamsRequestType) (*GetFeaturedStreamsResponseType, error) {
	q, err := query.Values(getFeaturedStreamsRequestType)
	if err != nil {
		return &GetFeaturedStreamsResponseType{}, err
	}
	u, err := url.Parse(session.URL + "/streams/featured?" + q.Encode())
	if err != nil {
		return &GetFeaturedStreamsResponseType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GetFeaturedStreamsResponseType{}, err
	}
	var s GetFeaturedStreamsResponseType
	err = json.Unmarshal([]byte(body), &s)
	if err != nil {
		return &GetFeaturedStreamsResponseType{}, err
	}

	return &s, nil
}

type GetStreamsSummaryRequestType struct {
	Game string `url:"game,omitempty"`
}

type GetStreamsSummaryResponseType struct {
	Viewers  int               `json:"viewers"`
	Links    map[string]string `json:"_links"`
	Channels int               `json:"channels"`
}

func (session *Session) GetStreamsSummary(getStreamsSummaryRequestType *GetStreamsSummaryRequestType) (*GetStreamsSummaryResponseType, error) {
	q, err := query.Values(getStreamsSummaryRequestType)
	if err != nil {
		return &GetStreamsSummaryResponseType{}, err
	}
	u, err := url.Parse(session.URL + "/streams/summary?" + q.Encode())
	if err != nil {
		return &GetStreamsSummaryResponseType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GetStreamsSummaryResponseType{}, err
	}
	var s GetStreamsSummaryResponseType
	err = json.Unmarshal([]byte(body), &s)
	if err != nil {
		return &GetStreamsSummaryResponseType{}, err
	}

	return &s, nil
}
