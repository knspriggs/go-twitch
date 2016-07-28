package twitch

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	clientID = os.Getenv("CLIENT_ID")
}

func TestGetTopGames(t *testing.T) {
	req := &GetTopGamesInputType{
		Limit:  10,
		Offset: 0,
	}
	session, err := NewSession(DefaultURL, APIV3Header, clientID)
	resp, err := session.GetTopGames(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, len(resp.Top), 10)
	}
}
