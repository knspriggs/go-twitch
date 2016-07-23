package twitch

//
// Generic streams types
//

// StreamType -
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

// FeaturedType -
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

// GetStreamsInputType -
type GetStreamsInputType struct {
	Game       string `url:"game,omitempty"`
	Channel    string `url:"channel,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
	ClientID   string `url:"client_id,omitempty"`
	StreamType string `url:"stream_type,omitempty"`
	Language   string `url:"language,omitempty"`
}

// GetStreamsOutputType -
type GetStreamsOutputType struct {
	Total   int               `json:"_total"`
	Streams []StreamType      `json:"streams"`
	Links   map[string]string `json:"_links"`
}

// GetStream -
func (session *Session) GetStream(getStreamsInputType *GetStreamsInputType) (*GetStreamsOutputType, error) {
	var out GetStreamsOutputType
	err := session.Request("GET", "/streams", &getStreamsInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetStreamByChannelInputType -
type GetStreamByChannelInputType struct {
	Channel string `url:"channel"`
}

// GetStreamByChannelOutputType -
type GetStreamByChannelOutputType struct {
	Stream StreamType        `json:"stream"`
	Links  map[string]string `json:"_links"`
}

// GetStreamByChannel -
func (session *Session) GetStreamByChannel(getStreamByChannelInputType *GetStreamByChannelInputType) (*GetStreamByChannelOutputType, error) {
	var out GetStreamByChannelOutputType
	err := session.Request("GET", "/streams/"+getStreamByChannelInputType.Channel, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetFeaturedStreamsInputType -
type GetFeaturedStreamsInputType struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// GetFeaturedStreamsOutputType -
type GetFeaturedStreamsOutputType struct {
	Featured []FeaturedType    `json:"featured"`
	Links    map[string]string `json:"_links"`
}

// GetFeaturedStreams -
func (session *Session) GetFeaturedStreams(getFeaturedStreamsInputType *GetFeaturedStreamsInputType) (*GetFeaturedStreamsOutputType, error) {
	var out GetFeaturedStreamsOutputType
	err := session.Request("GET", "/streams/featured", &getFeaturedStreamsInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetStreamsSummaryInputType -
type GetStreamsSummaryInputType struct {
	Game string `url:"game,omitempty"`
}

// GetStreamsSummaryOutputType -
type GetStreamsSummaryOutputType struct {
	Viewers  int               `json:"viewers"`
	Links    map[string]string `json:"_links"`
	Channels int               `json:"channels"`
}

// GetStreamsSummary -
func (session *Session) GetStreamsSummary(getStreamsSummaryInputType *GetStreamsSummaryInputType) (*GetStreamsSummaryOutputType, error) {
	var out GetStreamsSummaryOutputType
	err := session.Request("GET", "/streams/summary", &getStreamsSummaryInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
