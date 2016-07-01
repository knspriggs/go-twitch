package twitch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChannel(t *testing.T) {
	req := &GetChannelInputType{
		Channel: "Nightblue3",
	}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetChannel(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, resp.Views, 0)
	}
}
