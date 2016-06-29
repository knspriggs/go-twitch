package twitch

import (
  "testing"
)

func TestGetTopGames(t *testing.T) {
  grt := &GamesRequestType{
    Limit: 10,
    Offset: 0,
  }
  session, err := NewSession(DefaultURL, APIV3Header)
  gamesResponse, err := session.GetTopGames(grt)
  if err != nil {
    t.Fail()
  }
  if len(gamesResponse.Top) != 10 {
    t.Fail()
  }
}
