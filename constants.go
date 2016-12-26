package twitch

import "net/url"

const (
	// APIV3Header - default v3 api header
	APIV3Header = "application/vnd.twitchtv.v3+json"
)

var (
	// DefaultURL - default URLs for the Twitch v3 API
	DefaultURL = &url.URL{
		Scheme: "https",
		Host:   "api.twitch.tv",
		Path:   "kraken",
	}
)
