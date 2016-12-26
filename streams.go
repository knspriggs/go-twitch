package twitch

//
// Generic streams types
//

// StreamType - describes a stream
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

// FeaturedType - describes the relationship a stream has if it is featured
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

// GetStreamsInputType - request paramaters for the GetStream function
type GetStreamsInputType struct {
	Game       string `url:"game,omitempty"`
	Channel    string `url:"channel,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
	ClientID   string `url:"client_id,omitempty"`
	StreamType string `url:"stream_type,omitempty"`
	Language   string `url:"language,omitempty"`
}

// GetStreamsOutputType - response for the GetStream function
type GetStreamsOutputType struct {
	Total   int               `json:"_total"`
	Streams []StreamType      `json:"streams"`
	Links   map[string]string `json:"_links"`
}

// GetStream - returns the streams matching the input parameters
func (session *Session) GetStream(getStreamsInputType *GetStreamsInputType) (*GetStreamsOutputType, error) {
	var out GetStreamsOutputType
	err := session.request("GET", "/streams", &getStreamsInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetStreamByChannelInputType - specifies the channel for the GetStreamByChannel function
type GetStreamByChannelInputType struct {
	Channel string `url:"channel"`
}

// GetStreamByChannelOutputType - response for the GetStreamByChannel function
type GetStreamByChannelOutputType struct {
	Stream StreamType        `json:"stream"`
	Links  map[string]string `json:"_links"`
}

// GetStreamByChannel - returns the current stream for a channel
func (session *Session) GetStreamByChannel(getStreamByChannelInputType *GetStreamByChannelInputType) (*GetStreamByChannelOutputType, error) {
	var out GetStreamByChannelOutputType
	err := session.request("GET", "/streams/"+getStreamByChannelInputType.Channel, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetFeaturedStreamsInputType - input parameters for the GetFeaturedStreams function
type GetFeaturedStreamsInputType struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// GetFeaturedStreamsOutputType - contains a list of featured streams
type GetFeaturedStreamsOutputType struct {
	Featured []FeaturedType    `json:"featured"`
	Links    map[string]string `json:"_links"`
}

// GetFeaturedStreams - returns the featured streams
func (session *Session) GetFeaturedStreams(getFeaturedStreamsInputType *GetFeaturedStreamsInputType) (*GetFeaturedStreamsOutputType, error) {
	var out GetFeaturedStreamsOutputType
	err := session.request("GET", "/streams/featured", &getFeaturedStreamsInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetStreamsSummaryInputType - contains the game to scope the query to
type GetStreamsSummaryInputType struct {
	Game string `url:"game,omitempty"`
}

// GetStreamsSummaryOutputType - response object describing the summary of a game
type GetStreamsSummaryOutputType struct {
	Viewers  int               `json:"viewers"`
	Links    map[string]string `json:"_links"`
	Channels int               `json:"channels"`
}

// GetStreamsSummary - returns the summary of a game on twitch
func (session *Session) GetStreamsSummary(getStreamsSummaryInputType *GetStreamsSummaryInputType) (*GetStreamsSummaryOutputType, error) {
	var out GetStreamsSummaryOutputType
	err := session.request("GET", "/streams/summary", &getStreamsSummaryInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
