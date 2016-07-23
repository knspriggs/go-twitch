package twitch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopVideos(t *testing.T) {
	req := &GetTopVideosInputType{}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetTopVideos(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.True(t, len(resp.Videos) > 0)
	}
}

func TestGetChannelVideos(t *testing.T) {
	req := &GetChannelVideosInputType{
		Channel: "Nightblue3",
	}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetChannelVideos(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.True(t, len(resp.Videos) > 0)
	}
}
