package twitch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopGames(t *testing.T) {
	req := &GamesInputType{
		Limit:  10,
		Offset: 0,
	}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetTopGames(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, len(resp.Top), 10)
	}
}
