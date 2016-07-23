package twitch

// FollowsChannelType -
type FollowsChannelType struct {
	CreatedAt     string            `json:"created_at"`
	Notifications bool              `json:"notifications"`
	User          UserType          `json:"user"`
	Links         map[string]string `json:"_links"`
}

// FollowsUserType -
type FollowsUserType struct {
	CreatedAt     string            `json:"created_at"`
	Notifications bool              `json:"notifications"`
	Channel       ChannelType       `json:"channel"`
	Links         map[string]string `json:"_links"`
}

//
// Implementation and their respective request/response types
//

// GetChannelFollowsInputType -
type GetChannelFollowsInputType struct {
	Channel   string
	Limit     int    `url:"limit,omitempty"`
	Cursor    string `url:"cursor,omitempty"`
	Direction string `url:"direction,omitempty"`
}

// GetChannelFollowsOutputType -
type GetChannelFollowsOutputType struct {
	Total   int               `json:"_total"`
	Cursor  string            `json:"_cursor"`
	Follows []FollowsUserType `json:"follows"`
	Links   map[string]string `json:"_links"`
}

// GetChannelFollows -
func (session *Session) GetChannelFollows(getChannelFollowsInputType *GetChannelFollowsInputType) (*GetChannelFollowsOutputType, error) {
	var out GetChannelFollowsOutputType
	err := session.Request("GET", "/channels/"+getChannelFollowsInputType.Channel+"/follows", &getChannelFollowsInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetUserFollowsInputType -
type GetUserFollowsInputType struct {
	User      string
	Limit     int    `url:"limit,omitempty"`
	Direction string `url:"direction,omitempty"`
	SortyBy   string `url:"sortby,omitempty"`
}

// GetUserFollowsOutputType -
type GetUserFollowsOutputType struct {
	Total   int                  `json:"_total"`
	Cursor  string               `json:"_cursor"`
	Follows []FollowsChannelType `json:"follows"`
	Links   map[string]string    `json:"_links"`
}

// GetUserFollows -
func (session *Session) GetUserFollows(getUserFollowsInputType *GetUserFollowsInputType) (*GetUserFollowsOutputType, error) {
	var out GetUserFollowsOutputType
	err := session.Request("GET", "/users/"+getUserFollowsInputType.User+"/follows/channels", &getUserFollowsInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetUserFollowsChannelInputType -
type GetUserFollowsChannelInputType struct {
	User    string
	Channel string
}

// GetUserFollowsChannelOutputType -
type GetUserFollowsChannelOutputType struct {
	Follows       bool
	CreatedAt     string            `json:"created_at"`
	Notifications bool              `json:"notifications"`
	Channel       ChannelType       `json:"channel"`
	Links         map[string]string `json:"_links"`
}

// GetUserFollowsChannel -
func (session *Session) GetUserFollowsChannel(getUserFollowsChannelInputType *GetUserFollowsChannelInputType) (*GetUserFollowsChannelOutputType, error) {
	var out GetUserFollowsChannelOutputType
	err := session.Request("GET", "/users/"+getUserFollowsChannelInputType.User+"/follows/channels/"+getUserFollowsChannelInputType.Channel, nil, &out)
	if err != nil {
		return nil, err
	}
	out.Follows = true
	return &out, nil
}
