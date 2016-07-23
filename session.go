package twitch

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Session struct {
	Client        *http.Client
	URL           string
	VersionHeader string
}

func NewSession(url string, versionHeader string) (*Session, error) {
	client := &http.Client{}
	return &Session{
		Client:        client,
		URL:           url,
		VersionHeader: versionHeader,
	}, nil
}

func (session *Session) Request(method string, url string, r interface{}) error {
	request, requestError := http.NewRequest(method, url, bytes.NewBuffer([]byte("")))
	if requestError != nil {
		return requestError
	}
	request.Header.Add("Accept", APIV3Header)

	response, responseError := session.Client.Do(request)
	if responseError != nil {
		return responseError
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(body), r)
	if err != nil {
		return err
	}

	return nil
}
