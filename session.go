package twitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// NewSessionInput - input struct for creating a new session
type NewSessionInput struct {
	URL           *url.URL
	VersionHeader string
	ClientID      string // required
}

// Session represents a persistent connection to Twitch
type Session struct {
	Client        *http.Client
	URL           *url.URL
	VersionHeader string
	ClientID      string
}

type rootResponse struct {
	Links      map[string]string      `json:"links"`
	Identified bool                   `json:"identified"`
	Token      map[string]interface{} `json:"token"`
}

// NewSession creates and returns a new Twtich session
func NewSession(input NewSessionInput) (*Session, error) {
	if input.ClientID == "" {
		return nil, fmt.Errorf("A clientID must be supplied")
	}

	if input.URL == nil {
		input.URL = DefaultURL
	}

	if input.VersionHeader == "" {
		input.VersionHeader = APIV3Header
	}

	return &Session{
		Client:        &http.Client{},
		URL:           input.URL,
		VersionHeader: input.VersionHeader,
		ClientID:      input.ClientID,
	}, nil
}

func (session *Session) request(method string, url string, q interface{}, r interface{}) error {
	queryString, err := buildQueryString(q)
	request, requestError := http.NewRequest(method, session.URL.String()+url+queryString, bytes.NewBuffer([]byte("")))
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
	if err != nil {
		return err
	}
	response.Body.Close()

	if err := json.Unmarshal([]byte(body), r); err != nil {
		return err
	}

	return nil
}

// CheckClientID ensures that the client ID is correct.  This is done by
// performing a get to the root of twitch and confirming that the response's
// identified field is true
func (session *Session) CheckClientID() error {
	var rr rootResponse
	if err := session.request("GET", "/", nil, &rr); err != nil {
		return err
	}

	if !rr.Identified {
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
