package twitch

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/google/go-querystring/query"
)

//
// Implementation and their respective request/response types
//

type SearchChannelsInputType struct {
	Query  string `url:"query"`
	Limit  int    `url:"limit"`
	Offset int    `url:"offset"`
}

type SearchChannelsOutputType struct {
	Channels []ChannelType     `json:"channels"`
	Total    int               `json:"_total"`
	Links    map[string]string `json:"_links"`
}

func (session *Session) SearchChannels(searchChannelsInputType *SearchChannelsInputType) (*SearchChannelsOutputType, error) {
	q, err := query.Values(searchChannelsInputType)
	if err != nil {
		return &SearchChannelsOutputType{}, err
	}
	u, err := url.Parse(session.URL + "/search/channels?" + q.Encode())
	if err != nil {
		return &SearchChannelsOutputType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &SearchChannelsOutputType{}, err
	}
	var out SearchChannelsOutputType
	err = json.Unmarshal([]byte(body), &out)
	if err != nil {
		return &SearchChannelsOutputType{}, err
	}

	return &out, nil
}

type SearchStreamsInputType struct {
	Query  string `url:"query"`
	Limit  int    `url:"limit"`
	Offset int    `url:"offset"`
	HLS    bool   `url:"hls"`
}

type SearchStreamsOutputType struct {
	Streams []StreamType      `json:"streams"`
	Total   int               `json:"_total"`
	Links   map[string]string `json:"_links"`
}

func (session *Session) SearchStreams(searchStreamsInputType *SearchStreamsInputType) (*SearchStreamsOutputType, error) {
	q, err := query.Values(searchStreamsInputType)
	if err != nil {
		return &SearchStreamsOutputType{}, err
	}
	u, err := url.Parse(session.URL + "/search/streams?" + q.Encode())
	if err != nil {
		return &SearchStreamsOutputType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &SearchStreamsOutputType{}, err
	}
	var out SearchStreamsOutputType
	err = json.Unmarshal([]byte(body), &out)
	if err != nil {
		return &SearchStreamsOutputType{}, err
	}

	return &out, nil
}

type SearchGamesInputType struct {
	Query string `url:"query"`
	Type  string `url:"type"`
	Live  bool   `url:"live"`
}

type SearchGamesOutputType struct {
	Games []StreamType      `json:"games"`
	Links map[string]string `json:"_links"`
}

func (session *Session) SearchGames(searchGamesInputType *SearchGamesInputType) (*SearchGamesOutputType, error) {
	q, err := query.Values(searchGamesInputType)
	if err != nil {
		return &SearchGamesOutputType{}, err
	}
	u, err := url.Parse(session.URL + "/search/games?" + q.Encode())
	if err != nil {
		return &SearchGamesOutputType{}, err
	}

	res, _ := session.Request("GET", u.String())
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &SearchGamesOutputType{}, err
	}
	var out SearchGamesOutputType
	err = json.Unmarshal([]byte(body), &out)
	if err != nil {
		return &SearchGamesOutputType{}, err
	}

	return &out, nil
}
