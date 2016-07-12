package twitch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIngests(t *testing.T) {
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetIngests()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, len(resp.Ingests), 0)
	}
}
