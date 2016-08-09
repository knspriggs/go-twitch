package twitch_test

import "net/url"

var (
	DefaultURL = &url.URL{
		Scheme: "https",
		Host:   "api.twitch.tv",
		Path:   "kraken",
	}
)
