package twitch

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Session -
type Session struct {
	Client        *http.Client
	URL           string
	VersionHeader string
}

// NewSession -
func NewSession(url string, versionHeader string) (*Session, error) {
	client := &http.Client{}
	return &Session{
		Client:        client,
		URL:           url,
		VersionHeader: versionHeader,
	}, nil
}

func (session *Session) request(method string, url string, q interface{}, r interface{}) error {
	var queryString string
	if q != nil {
		query, err := query.Values(q)
		if err != nil {
			return err
		}
		queryString = "?" + query.Encode()
	} else {
		queryString = ""
	}
	request, requestError := http.NewRequest(method, session.URL+url+queryString, bytes.NewBuffer([]byte("")))
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
