package twitch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChannelFollows(t *testing.T) {
	req := &GetChannelFollowsInputType{
		Channel: "knspriggs",
	}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetChannelFollows(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, resp.Total, 0)
	}
}

func TestGetUserFollows(t *testing.T) {
	req := &GetUserFollowsInputType{
		User: "knspriggs",
	}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetUserFollows(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, resp.Total, 0)
	}
}

func TestGetUserFollowsChannelValid(t *testing.T) {
	req := &GetUserFollowsChannelInputType{
		User:    "knspriggs",
		Channel: "nightblue3",
	}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetUserFollowsChannel(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.True(t, resp.Follows)
	}
}

func TestGetUserFollowsChannelInvalid(t *testing.T) {
	req := &GetUserFollowsChannelInputType{
		User:    "knspriggs",
		Channel: "sdfknaosfg",
	}
	session, err := NewSession(DefaultURL, APIV3Header)
	resp, err := session.GetUserFollowsChannel(req)
	assert.Nil(t, err)
	assert.False(t, resp.Follows)
}
