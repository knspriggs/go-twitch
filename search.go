package twitch

//
// Implementation and their respective request/response types
//

// SearchChannelsInputType - specifies the query parameters for a channel search
type SearchChannelsInputType struct {
	Query  string `url:"query,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

// SearchChannelsOutputType - contains the results for a channel search
type SearchChannelsOutputType struct {
	Channels []ChannelType     `json:"channels"`
	Total    int               `json:"_total"`
	Links    map[string]string `json:"_links"`
}

// SearchChannels - returns channels matching the query
func (session *Session) SearchChannels(searchChannelsInputType *SearchChannelsInputType) (*SearchChannelsOutputType, error) {
	var out SearchChannelsOutputType
	err := session.request("GET", "/search/channels", &searchChannelsInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// SearchStreamsInputType - specifies the query parameters for a stream search
type SearchStreamsInputType struct {
	Query  string `url:"query,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
	HLS    bool   `url:"hls,omitempty"`
}

// SearchStreamsOutputType - contains the results for a stream search
type SearchStreamsOutputType struct {
	Streams []StreamType      `json:"streams"`
	Total   int               `json:"_total"`
	Links   map[string]string `json:"_links"`
}

// SearchStreams - returns streams matching the query
func (session *Session) SearchStreams(searchStreamsInputType *SearchStreamsInputType) (*SearchStreamsOutputType, error) {
	var out SearchStreamsOutputType
	err := session.request("GET", "/search/streams", &searchStreamsInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// SearchGamesInputType - specifies the query parameters for a game search
type SearchGamesInputType struct {
	Query string `url:"query,omitempty"`
	Type  string `url:"type,omitempty"`
	Live  bool   `url:"live,omitempty"`
}

// SearchGamesOutputType - contains the results for a game search
type SearchGamesOutputType struct {
	Games []StreamType      `json:"games"`
	Links map[string]string `json:"_links"`
}

// SearchGames - returns games matching the query
func (session *Session) SearchGames(searchGamesInputType *SearchGamesInputType) (*SearchGamesOutputType, error) {
	var out SearchGamesOutputType
	err := session.request("GET", "/search/games", &searchGamesInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
