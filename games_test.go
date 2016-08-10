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

func TestGetTopGames(t *testing.T) {
	req := &twitch.GetTopGamesInputType{
		Limit:  10,
		Offset: 0,
	}
	session, err := twitch.NewSession(twitch.NewSessionInput{ClientID: clientID})
	resp, err := session.GetTopGames(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, len(resp.Top), 10)
	}
}
