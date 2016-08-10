package twitch

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	clientID = os.Getenv("CLIENT_ID")
}

func TestGetAllTeams(t *testing.T) {
	session, err := NewSession(NewSessionInput{ClientID: clientID})
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
	session, err := NewSession(NewSessionInput{ClientID: clientID})
	resp, err := session.GetTeam(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.True(t, resp.ID > 0)
	}
}
