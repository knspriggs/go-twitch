package twitch

import "net/url"

const (
	// APIV3Header -
	APIV3Header = "application/vnd.twitchtv.v3+json"
)

var (
	DefaultURL = &url.URL{
		Scheme: "https",
		Host:   "api.twitch.tv",
		Path:   "kraken",
	}
)
