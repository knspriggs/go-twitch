package twitch

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	clientID = os.Getenv("CLIENT_ID")
}

func TestGetIngests(t *testing.T) {
	session, err := NewSession(DefaultURL, APIV3Header, clientID)
	resp, err := session.GetIngests()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, len(resp.Ingests), 0)
	}
}
