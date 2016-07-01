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

type GetStreamsInputType struct {
	Game       string `url:"game,omitempty"`
	Channel    string `url:"channel,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
	ClientID   string `url:"client_id,omitempty"`
	StreamType string `url:"stream_type,omitempty"`
	Language   string `url:"language,omitempty"`
}

type GetStreamsOutputType struct {
	Total   int               `json:"_total"`
	Streams []StreamType      `json:"streams"`
	Links   map[string]string `json:"_links"`
}

func (session *Session) GetStream(getStreamsInputType *GetStreamsInputType) (*GetStreamsOutputType, error) {
	q, err := query.Values(getStreamsInputType)
	if err != nil {
		return &GetStreamsOutputType{}, err
	}
	u, err := url.Parse(session.URL + "/streams?" + q.Encode())
	if err != nil {
		return &GetStreamsOutputType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GetStreamsOutputType{}, err
	}
	var s GetStreamsOutputType
	err = json.Unmarshal([]byte(body), &s)
	if err != nil {
		return &GetStreamsOutputType{}, err
	}

	return &s, nil
}

type GetStreamByChannelInputType struct {
	Channel string `url:"channel"`
}

type GetStreamByChannelOutputType struct {
	Stream StreamType        `json:"stream"`
	Links  map[string]string `json:"_links"`
}

func (session *Session) GetStreamByChannel(getStreamByChannelInputType *GetStreamByChannelInputType) (*GetStreamByChannelOutputType, error) {
	u, err := url.Parse(session.URL + "/streams/" + getStreamByChannelInputType.Channel)
	if err != nil {
		return &GetStreamByChannelOutputType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GetStreamByChannelOutputType{}, err
	}
	var sbc GetStreamByChannelOutputType
	err = json.Unmarshal([]byte(body), &sbc)
	if err != nil {
		return &GetStreamByChannelOutputType{}, err
	}

	return &sbc, nil
}

type GetFeaturedStreamsInputType struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type GetFeaturedStreamsOutputType struct {
	Featured []FeaturedType    `json:"featured"`
	Links    map[string]string `json:"_links"`
}

func (session *Session) GetFeaturedStreams(getFeaturedStreamsInputType *GetFeaturedStreamsInputType) (*GetFeaturedStreamsOutputType, error) {
	q, err := query.Values(getFeaturedStreamsInputType)
	if err != nil {
		return &GetFeaturedStreamsOutputType{}, err
	}
	u, err := url.Parse(session.URL + "/streams/featured?" + q.Encode())
	if err != nil {
		return &GetFeaturedStreamsOutputType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GetFeaturedStreamsOutputType{}, err
	}
	var s GetFeaturedStreamsOutputType
	err = json.Unmarshal([]byte(body), &s)
	if err != nil {
		return &GetFeaturedStreamsOutputType{}, err
	}

	return &s, nil
}

type GetStreamsSummaryInputType struct {
	Game string `url:"game,omitempty"`
}

type GetStreamsSummaryOutputType struct {
	Viewers  int               `json:"viewers"`
	Links    map[string]string `json:"_links"`
	Channels int               `json:"channels"`
}

func (session *Session) GetStreamsSummary(getStreamsSummaryInputType *GetStreamsSummaryInputType) (*GetStreamsSummaryOutputType, error) {
	q, err := query.Values(getStreamsSummaryInputType)
	if err != nil {
		return &GetStreamsSummaryOutputType{}, err
	}
	u, err := url.Parse(session.URL + "/streams/summary?" + q.Encode())
	if err != nil {
		return &GetStreamsSummaryOutputType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &GetStreamsSummaryOutputType{}, err
	}
	var s GetStreamsSummaryOutputType
	err = json.Unmarshal([]byte(body), &s)
	if err != nil {
		return &GetStreamsSummaryOutputType{}, err
	}

	return &s, nil
}
