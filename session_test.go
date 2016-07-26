package twitch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
