package twitch

import (
  "testing"
)

func TestGetStreamByChannel(t *testing.T) {
  gst := &GetStreamByChannelRequestType{
    Channel: "#knspriggs",
  }
  session, err := NewSession(DefaultURL, APIV3Header)
  getStreamByChannelResponse, err := session.GetStreamByChannel(gst)
  if err != nil {
    t.Logf("got error: %s", err.Error())
    t.Fail()
  }
  if getStreamByChannelResponse.Stream.Viewers != 0 {
    t.Fail()
  }
}

func TestGetStreamsWithoutRequestParams(t *testing.T) {
  gst := &GetStreamsRequestType{}
  session, err := NewSession(DefaultURL, APIV3Header)
  getStreamsResponse, err := session.GetStream(gst)
  if err != nil {
    t.Logf("got error: %s", err.Error())
    t.Fail()
  }
  if len(getStreamsResponse.Streams) == 0 {
    t.Fail()
  }
}

func TestGetStreamsWithPartialRequestParamsAndDefaults(t *testing.T) {
  gst := &GetStreamsRequestType{
    Game: "Counter-Strike: Global Offensive",
  }
  session, err := NewSession(DefaultURL, APIV3Header)
  getStreamsResponse, err := session.GetStream(gst)
  if err != nil {
    t.Logf("got error: %s", err.Error())
    t.Fail()
  }
  if len(getStreamsResponse.Streams) != 25 {
    t.Fail()
  }
}

func TestGetStreamsWithPartialRequestParams(t *testing.T) {
  gst := &GetStreamsRequestType{
    Game: "Counter-Strike: Global Offensive",
    Limit: 10,
    Offset: 1,
    StreamType: "live",
  }
  session, err := NewSession(DefaultURL, APIV3Header)
  getStreamsResponse, err := session.GetStream(gst)
  if err != nil {
    t.Logf("got error: %s", err.Error())
    t.Fail()
  }
  if len(getStreamsResponse.Streams) != 10 {
    t.Fail()
  }
}
