package twitch

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	clientID = os.Getenv("CLIENT_ID")
}

func TestGetTopVideos(t *testing.T) {
	req := &GetTopVideosInputType{}
	session, err := NewSession(NewSessionInput{ClientID: clientID})
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
	session, err := NewSession(NewSessionInput{ClientID: clientID})
	resp, err := session.GetChannelVideos(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.True(t, len(resp.Videos) > 0)
	}
}
