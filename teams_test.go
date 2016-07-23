package twitch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTeams(t *testing.T) {
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetAllTeams()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.True(t, len(resp.Teams) > 0)
	}
}

func TestGetTeam(t *testing.T) {
	req := &GetTeamInputType{
		Team: "tckt",
	}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetTeam(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.True(t, resp.ID > 0)
	}
}
