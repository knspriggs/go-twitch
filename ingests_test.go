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

func TestGetIngests(t *testing.T) {
	session, err := twitch.NewSession(twitch.NewSessionInput{ClientID: clientID})
	resp, err := session.GetIngests()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, len(resp.Ingests), 0)
	}
}
