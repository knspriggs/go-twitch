package twitch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStreamByChannel(t *testing.T) {
	req := &GetStreamByChannelInputType{
		Channel: "#knspriggs",
	}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetStreamByChannel(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Stream.Viewers, 0)
	}
}

func TestGetStreamsWithoutRequestParams(t *testing.T) {
	req := &GetStreamsInputType{}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetStream(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, resp.Streams, 0)
	}
}

func TestGetStreamsWithPartialRequestParamsAndDefaults(t *testing.T) {
	req := &GetStreamsInputType{
		Game: "Counter-Strike: Global Offensive",
	}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetStream(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, len(resp.Streams), 25)
	}
}

func TestGetStreamsWithPartialRequestParams(t *testing.T) {
	req := &GetStreamsInputType{
		Game:       "Counter-Strike: Global Offensive",
		Limit:      10,
		Offset:     1,
		StreamType: "live",
	}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetStream(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, len(resp.Streams), 10)
	}
}

func TestGetFeaturedStreams(t *testing.T) {
	req := &GetFeaturedStreamsInputType{}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetFeaturedStreams(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, len(resp.Featured), 25)
	}
}

func TestGetStreamsSummaryWithGame(t *testing.T) {
	req := &GetStreamsSummaryInputType{
		Game: "Counter-Strike: Global Offensive",
	}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetStreamsSummary(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, resp.Viewers, 0)
	}
}

func TestGetStreamsSummaryWithoutGame(t *testing.T) {
	req := &GetStreamsSummaryInputType{}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetStreamsSummary(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, resp.Viewers, 0)
	}
}
