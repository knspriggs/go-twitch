package twitch

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

type FollowsChannelType struct {
	CreatedAt     string            `json:"created_at"`
	Notifications bool              `json:"notifications"`
	User          UserType          `json:"user"`
	Links         map[string]string `json:"_links"`
}

type FollowsUserType struct {
	CreatedAt     string            `json:"created_at"`
	Notifications bool              `json:"notifications"`
	Channel       ChannelType       `json:"channel"`
	Links         map[string]string `json:"_links"`
}

//
// Implementation and their respective request/response types
//

type GetChannelFollowsInputType struct {
	Channel   string
	Limit     int    `url:"limit"`
	Cursor    string `url:"cursor"`
	Direction string `url:"direction"`
}
type GetChannelFollowsOutputType struct {
	Total   int               `json:"_total"`
	Cursor  string            `json:"_cursor"`
	Follows []FollowsUserType `json:"follows"`
	Links   map[string]string `json:"_links"`
}

func (session *Session) GetChannelFollows(getChannelFollowsInputType *GetChannelFollowsInputType) (*GetChannelFollowsOutputType, error) {
	q, err := query.Values(getChannelFollowsInputType)
	if err != nil {
		return &GetChannelFollowsOutputType{}, err
	}
	u, err := url.Parse(session.URL + "/channels/" + getChannelFollowsInputType.Channel + "/follows?" + q.Encode())
	if err != nil {
		return &GetChannelFollowsOutputType{}, err
	}

	var out GetChannelFollowsOutputType
	err = session.Request("GET", u.String(), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type GetUserFollowsInputType struct {
	User      string
	Limit     int    `url:"limit"`
	Direction string `url:"direction"`
	SortyBy   string `url:"sortby"`
}
type GetUserFollowsOutputType struct {
	Total   int                  `json:"_total"`
	Cursor  string               `json:"_cursor"`
	Follows []FollowsChannelType `json:"follows"`
	Links   map[string]string    `json:"_links"`
}

func (session *Session) GetUserFollows(getUserFollowsInputType *GetUserFollowsInputType) (*GetUserFollowsOutputType, error) {
	q, err := query.Values(getUserFollowsInputType)
	if err != nil {
		return &GetUserFollowsOutputType{}, err
	}
	u, err := url.Parse(session.URL + "/users/" + getUserFollowsInputType.User + "/follows/channels?" + q.Encode())
	if err != nil {
		return &GetUserFollowsOutputType{}, err
	}

	var out GetUserFollowsOutputType
	err = session.Request("GET", u.String(), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type GetUserFollowsChannelInputType struct {
	User    string
	Channel string
}
type GetUserFollowsChannelOutputType struct {
	Follows       bool
	CreatedAt     string            `json:"created_at"`
	Notifications bool              `json:"notifications"`
	Channel       ChannelType       `json:"channel"`
	Links         map[string]string `json:"_links"`
}

func (session *Session) GetUserFollowsChannel(getUserFollowsChannelInputType *GetUserFollowsChannelInputType) (*GetUserFollowsChannelOutputType, error) {
	u, err := url.Parse(session.URL + "/users/" + getUserFollowsChannelInputType.User + "/follows/channels/" + getUserFollowsChannelInputType.Channel)
	if err != nil {
		return &GetUserFollowsChannelOutputType{}, err
	}

	var out GetUserFollowsChannelOutputType
	err = session.Request("GET", u.String(), &out)
	if err != nil {
		return nil, err
	}
	out.Follows = true
	return &out, nil
}
