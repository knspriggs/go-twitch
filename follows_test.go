package twitch_test

import (
	"os"
	"testing"

	twitch "github.com/knspriggs/go-twitch"
	"github.com/stretchr/testify/assert"
)

func init() {
	clientID = os.Getenv("CLIENT_ID")
}

func TestGetChannelFollows(t *testing.T) {
	req := &twitch.GetChannelFollowsInputType{
		Channel: "knspriggs",
	}
	session, err := twitch.NewSession(twitch.NewSessionInput{ClientID: clientID})
	resp, err := session.GetChannelFollows(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, resp.Total, 0)
	}
}

func TestGetUserFollows(t *testing.T) {
	req := &twitch.GetUserFollowsInputType{
		User: "knspriggs",
	}
	session, err := twitch.NewSession(twitch.NewSessionInput{ClientID: clientID})
	resp, err := session.GetUserFollows(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, resp.Total, 0)
	}
}

func TestGetUserFollowsChannelValid(t *testing.T) {
	req := &twitch.GetUserFollowsChannelInputType{
		User:    "knspriggs",
		Channel: "nightblue3",
	}
	session, err := twitch.NewSession(twitch.NewSessionInput{ClientID: clientID})
	resp, err := session.GetUserFollowsChannel(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.True(t, resp.Follows)
	}
}

func TestGetUserFollowsChannelInvalid(t *testing.T) {
	t.Skip()
	req := &twitch.GetUserFollowsChannelInputType{
		User:    "knspriggs",
		Channel: "sdfknaosfg",
	}
	session, err := twitch.NewSession(twitch.NewSessionInput{ClientID: clientID})
	resp, err := session.GetUserFollowsChannel(req)
	assert.Nil(t, err)
	assert.False(t, resp.Follows)
}
