package twitch

//
// Implementation and their respective request/response types
//

// SearchChannelsInputType -
type SearchChannelsInputType struct {
	Query  string `url:"query,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

// SearchChannelsOutputType -
type SearchChannelsOutputType struct {
	Channels []ChannelType     `json:"channels"`
	Total    int               `json:"_total"`
	Links    map[string]string `json:"_links"`
}

// SearchChannels -
func (session *Session) SearchChannels(searchChannelsInputType *SearchChannelsInputType) (*SearchChannelsOutputType, error) {
	var out SearchChannelsOutputType
	err := session.Request("GET", "/search/channels", &searchChannelsInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// SearchStreamsInputType -
type SearchStreamsInputType struct {
	Query  string `url:"query,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
	HLS    bool   `url:"hls,omitempty"`
}

// SearchStreamsOutputType -
type SearchStreamsOutputType struct {
	Streams []StreamType      `json:"streams"`
	Total   int               `json:"_total"`
	Links   map[string]string `json:"_links"`
}

// SearchStreams -
func (session *Session) SearchStreams(searchStreamsInputType *SearchStreamsInputType) (*SearchStreamsOutputType, error) {
	var out SearchStreamsOutputType
	err := session.Request("GET", "/search/streams", &searchStreamsInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// SearchGamesInputType -
type SearchGamesInputType struct {
	Query string `url:"query,omitempty"`
	Type  string `url:"type,omitempty"`
	Live  bool   `url:"live,omitempty"`
}

// SearchGamesOutputType -
type SearchGamesOutputType struct {
	Games []StreamType      `json:"games"`
	Links map[string]string `json:"_links"`
}

// SearchGames -
func (session *Session) SearchGames(searchGamesInputType *SearchGamesInputType) (*SearchGamesOutputType, error) {
	var out SearchGamesOutputType
	err := session.Request("GET", "/search/games", &searchGamesInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
