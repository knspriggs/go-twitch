package twitch

import (
  "net/http"
  "bytes"
)

type Session struct {
  Client *http.Client
  URL string
  VersionHeader string
}

func NewSession(url string, versionHeader string) (*Session, error) {
  client := &http.Client{}
  return &Session{
    Client: client,
    URL: url,
    VersionHeader: versionHeader,
  }, nil
}

func (session *Session) Request(method string, url string) (*http.Response, error) {
  request, requestError := http.NewRequest(method, url, bytes.NewBuffer([]byte("")))
  if requestError != nil {
    return &http.Response{}, requestError
  }
  request.Header.Add("Accept", APIV3Header)

  response, responseError := session.Client.Do(request)
  if responseError != nil {
    return &http.Response{}, responseError
  }
  return response, nil
}
