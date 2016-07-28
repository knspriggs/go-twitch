package twitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Session -
type Session struct {
	Client        *http.Client
	URL           string
	VersionHeader string
	ClientID      string
}

// NewSession -
func NewSession(url string, versionHeader string, clientID string) (*Session, error) {
	client := &http.Client{}
	return &Session{
		Client:        client,
		URL:           url,
		VersionHeader: versionHeader,
		ClientID:      clientID,
	}, nil
}

func (session *Session) request(method string, url string, q interface{}, r interface{}) error {
	queryString, err := buildQueryString(q)
	request, requestError := http.NewRequest(method, session.URL+url+queryString, bytes.NewBuffer([]byte("")))
	if requestError != nil {
		return requestError
	}
	request.Header.Add("Accept", APIV3Header)
	request.Header.Add("Client-ID", session.ClientID)

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

type rootResponseType struct {
	Links      map[string]string      `json:"links"`
	Identified bool                   `json:"identified"`
	Token      map[string]interface{} `json:"token"`
}

// CheckClientID -
func (session *Session) CheckClientID() error {
	var rrt rootResponseType
	err := session.request("GET", "/", nil, &rrt)
	if err != nil {
		return err
	}
	if rrt.Identified != true {
		return fmt.Errorf("Session not identified, please check your client-id")
	}
	return nil
}

func buildQueryString(q interface{}) (string, error) {
	if q != nil {
		query, err := query.Values(q)
		if err != nil {
			return "", err
		}
		return "?" + query.Encode(), nil
	}
	return "", nil
}
