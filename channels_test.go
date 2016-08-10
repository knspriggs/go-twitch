package twitch_test

import (
	"os"
	"testing"

	twitch "github.com/knspriggs/go-twitch"
	"github.com/stretchr/testify/assert"
)

var clientID string

func init() {
	clientID = os.Getenv("CLIENT_ID")
}

func TestGetChannel(t *testing.T) {
	req := &twitch.GetChannelInputType{
		Channel: "Nightblue3",
	}
	session, err := twitch.NewSession(twitch.NewSessionInput{ClientID: clientID})
	resp, err := session.GetChannel(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, resp.Views, 0)
	}
}

func TestGetChannelTeams(t *testing.T) {
	req := &twitch.GetChannelTeamsInputType{
		Channel: "Nightblue3",
	}
	session, err := twitch.NewSession(twitch.NewSessionInput{ClientID: clientID})
	resp, err := session.GetChannelTeams(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, len(resp.Teams), 0)
	}
}
