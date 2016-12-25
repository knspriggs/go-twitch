[![Build Status](https://travis-ci.com/knspriggs/go-twitch.svg?token=zZCoL2DxeY3FuDqHfbp7&branch=add-travisci-yaml)](https://travis-ci.com/knspriggs/go-twitch)

# go-twitch


### Test
```
CLIENT_ID="<my client ID>" go test -v -cover
```

### Usage

Example File:
```go
package main

import (
  "log"
  "os"

  "github.com/knspriggs/go-twitch"
)

var clientID string

func init() {
  clientID = os.Getenv("CLIENT_ID")
}

func main() {
  twitchSession, err := twitch.NewSession(twitch.NewSessionInput{ClientID: clientID})
  if err != nil {
    log.Fatalln(err)
  }

  searchChannelsInput := twitch.SearchChannelsInputType{
    Query: "knspriggs",   // see https://github.com/justintv/Twitch-API/blob/master/v3_resources/search.md for query syntax
    Limit: 2,             // optional
    Offset: 0,            // optional
  }

  resp, err := twitchSession.SearchChannels(&searchChannelsInput)
  if err != nil {
    log.Fatalln(err)
  }
  log.Printf("Resp: \n%#v", resp)
}
```

```
CLIENT_ID="<my client ID>" go run example/main.go
```

To get a client ID see the documentation from the Twitch API https://github.com/justintv/Twitch-API/blob/master/authentication.md
