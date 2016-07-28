package twitch

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	clientID = os.Getenv("CLIENT_ID")
}

func TestSearchChannels(t *testing.T) {
	req := &SearchChannelsInputType{
		Query: "knspriggs",
	}
	session, err := NewSession(DefaultURL, APIV3Header, clientID)
	resp, err := session.SearchChannels(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Total, 1)
	}
}

func TestSearchChannelsEmpty(t *testing.T) {
	req := &SearchChannelsInputType{
		Query: "agjansa",
	}
	session, err := NewSession(DefaultURL, APIV3Header, clientID)
	resp, err := session.SearchChannels(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Total, 0)
	}
}

func TestSearchStreams(t *testing.T) {
	req := &SearchStreamsInputType{
		Query: "League",
	}
	session, err := NewSession(DefaultURL, APIV3Header, clientID)
	resp, err := session.SearchStreams(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, resp.Total, 0)
	}
}

func TestSearchStreamsEmpty(t *testing.T) {
	req := &SearchStreamsInputType{
		Query: "agjansa",
	}
	session, err := NewSession(DefaultURL, APIV3Header, clientID)
	resp, err := session.SearchStreams(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Total, 0)
	}
}

func TestSearchGames(t *testing.T) {
	req := &SearchGamesInputType{
		Query: "League",
		Type:  "suggest",
		Live:  true,
	}
	session, err := NewSession(DefaultURL, APIV3Header, clientID)
	resp, err := session.SearchGames(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, len(resp.Games), 0)
	}
}

func TestSearchGamesEmpty(t *testing.T) {
	req := &SearchGamesInputType{
		Query: "agjansa",
	}
	session, err := NewSession(DefaultURL, APIV3Header, clientID)
	resp, err := session.SearchGames(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, len(resp.Games), 0)
	}
}
