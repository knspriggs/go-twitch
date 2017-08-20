package twitch

import "net/url"

const (
	// APIV5Header - default v3 api header
	APIV5Header = "application/vnd.twitchtv.v5+json"
)

var (
	// DefaultURL - default URLs for the Twitch v3 API
	DefaultURL = &url.URL{
		Scheme: "https",
		Host:   "api.twitch.tv",
		Path:   "kraken",
	}
)
