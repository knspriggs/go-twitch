package twitch

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

//
// Generic streams types
//

// VideoType -
type VideoType struct {
	Title         string            `json:"title"`
	Description   string            `json:"description"`
	BroadcastID   int64             `json:"broadcast_id"`
	Status        string            `json:"status"`
	ID            string            `json:"_id"`
	TagList       string            `json:"tag_list"`
	RecordedAt    string            `json:"recorded_at"`
	Game          interface{}       `json:"game"`
	Length        int               `json:"length"`
	Preview       string            `json:"preview"`
	URL           string            `json:"url"`
	Views         int               `json:"views"`
	BroadcastType string            `json:"broadcast_type"`
	Links         map[string]string `json:"_links"`
	Channel       struct {
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
	} `json:"channel"`
}

//
// Implementation and their respective request/response types
//

// GetTopVideosInputType -
type GetTopVideosInputType struct {
	Game   string `url:"game,omitempty"`
	Period string `url:"period,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

// GetTopVideosOutputType -
type GetTopVideosOutputType struct {
	Videos []VideoType       `json:"videos"`
	Links  map[string]string `json:"_links"`
}

// GetTopVideos -
func (session *Session) GetTopVideos(getTopVideosInputType *GetTopVideosInputType) (*GetTopVideosOutputType, error) {
	q, err := query.Values(getTopVideosInputType)
	if err != nil {
		return &GetTopVideosOutputType{}, err
	}
	u, err := url.Parse(session.URL + "/videos/top?" + q.Encode())
	if err != nil {
		return nil, err
	}

	var out GetTopVideosOutputType
	err = session.Request("GET", u.String(), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetChannelVideosInputType -
type GetChannelVideosInputType struct {
	Channel    string
	Broadcasts bool `url:"broadcasts,omitempty"`
	HLS        bool `url:"hls,omitempty"`
	Limit      int  `url:"limit,omitempty"`
	Offset     int  `url:"offset,omitempty"`
}

// GetChannelVideosOutputType -
type GetChannelVideosOutputType struct {
	Videos []VideoType       `json:"videos"`
	Total  int               `json:"total"`
	Links  map[string]string `json:"_links"`
}

// GetChannelVideos -
func (session *Session) GetChannelVideos(getChannelVideosInputType *GetChannelVideosInputType) (*GetChannelVideosOutputType, error) {
	q, err := query.Values(getChannelVideosInputType)
	if err != nil {
		return &GetChannelVideosOutputType{}, err
	}
	u, err := url.Parse(session.URL + "/channels/" + getChannelVideosInputType.Channel + "/videos?" + q.Encode())
	if err != nil {
		return nil, err
	}

	var out GetChannelVideosOutputType
	err = session.Request("GET", u.String(), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
