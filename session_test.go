package twitch

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var clientID string
var DefaultURL *url.URL

func init() {
	clientID = os.Getenv("CLIENT_ID")
	DefaultURL = &url.URL{
		Scheme: "https",
		Host:   "api.twitch.tv",
		Path:   "kraken",
	}
}

func TestNewSession(t *testing.T) {
	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		t.Skip()
	}

	session, err := NewSession(DefaultURL, APIV3Header, clientID)
	err = session.CheckClientID()
	assert.Nil(t, err)
}

type queryTest struct {
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
	Game   string `url:"game,omitempty"`
}

func TestBuildQueryStringPartial(t *testing.T) {
	query := &queryTest{
		Limit: 10,
		Game:  "Destiny",
	}
	queryString, err := buildQueryString(query)
	assert.Nil(t, err)
	assert.Equal(t, "?game=Destiny&limit=10", queryString)
}

func TestBuildQueryStringEmpty(t *testing.T) {
	query := &queryTest{}
	queryString, err := buildQueryString(query)
	assert.Nil(t, err)
	assert.Equal(t, "?", queryString)
}
