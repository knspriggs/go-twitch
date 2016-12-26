package twitch

// FollowsChannelType - struct describing a following channel
type FollowsChannelType struct {
	CreatedAt     string            `json:"created_at"`
	Notifications bool              `json:"notifications"`
	User          UserType          `json:"user"`
	Links         map[string]string `json:"_links"`
}

// FollowsUserType - struct describing a following user
type FollowsUserType struct {
	CreatedAt     string            `json:"created_at"`
	Notifications bool              `json:"notifications"`
	Channel       ChannelType       `json:"channel"`
	Links         map[string]string `json:"_links"`
}

//
// Implementation and their respective request/response types
//

// GetChannelFollowsInputType - request type for querying users following a channel
type GetChannelFollowsInputType struct {
	Channel   string
	Limit     int    `url:"limit,omitempty"`
	Cursor    string `url:"cursor,omitempty"`
	Direction string `url:"direction,omitempty"`
}

// GetChannelFollowsOutputType - response type containing users following a channel
type GetChannelFollowsOutputType struct {
	Total   int               `json:"_total"`
	Cursor  string            `json:"_cursor"`
	Follows []FollowsUserType `json:"follows"`
	Links   map[string]string `json:"_links"`
}

// GetChannelFollows - returns the users who follow the specified channel
func (session *Session) GetChannelFollows(getChannelFollowsInputType *GetChannelFollowsInputType) (*GetChannelFollowsOutputType, error) {
	var out GetChannelFollowsOutputType
	err := session.request("GET", "/channels/"+getChannelFollowsInputType.Channel+"/follows", &getChannelFollowsInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetUserFollowsInputType - request type for querying what channels a user follows
type GetUserFollowsInputType struct {
	User      string
	Limit     int    `url:"limit,omitempty"`
	Direction string `url:"direction,omitempty"`
	SortyBy   string `url:"sortby,omitempty"`
}

// GetUserFollowsOutputType - response type containing channels a user follows
type GetUserFollowsOutputType struct {
	Total   int                  `json:"_total"`
	Cursor  string               `json:"_cursor"`
	Follows []FollowsChannelType `json:"follows"`
	Links   map[string]string    `json:"_links"`
}

// GetUserFollows - returns the channels that a specified user follows
func (session *Session) GetUserFollows(getUserFollowsInputType *GetUserFollowsInputType) (*GetUserFollowsOutputType, error) {
	var out GetUserFollowsOutputType
	err := session.request("GET", "/users/"+getUserFollowsInputType.User+"/follows/channels", &getUserFollowsInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetUserFollowsChannelInputType - request type container user and channel
type GetUserFollowsChannelInputType struct {
	User    string
	Channel string
}

// GetUserFollowsChannelOutputType - response type container wether a user follows the spcified channel
type GetUserFollowsChannelOutputType struct {
	Follows       bool
	CreatedAt     string            `json:"created_at"`
	Notifications bool              `json:"notifications"`
	Channel       ChannelType       `json:"channel"`
	Links         map[string]string `json:"_links"`
}

// GetUserFollowsChannel - returns wether a user follows the specified channel
func (session *Session) GetUserFollowsChannel(getUserFollowsChannelInputType *GetUserFollowsChannelInputType) (*GetUserFollowsChannelOutputType, error) {
	var out GetUserFollowsChannelOutputType
	err := session.request("GET", "/users/"+getUserFollowsChannelInputType.User+"/follows/channels/"+getUserFollowsChannelInputType.Channel, nil, &out)
	if err != nil {
		return nil, err
	}
	out.Follows = true
	return &out, nil
}
