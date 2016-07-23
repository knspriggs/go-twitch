package twitch

import (
	"net/url"
)

type IngestType struct {
	Name         string  `json:"name"`
	Default      bool    `json:"default"`
	ID           int     `json:"_id"`
	URLTemplate  string  `json:"url_template"`
	Availability float64 `json:"availability"`
}

//
// Implementation and their respective request/response types
//

type GetIngestsOutputType struct {
	Ingests []IngestType      `json:"ingests"`
	Links   map[string]string `json:"_links"`
}

func (session *Session) GetIngests() (*GetIngestsOutputType, error) {
	u, err := url.Parse(session.URL + "/ingests")
	if err != nil {
		return &GetIngestsOutputType{}, err
	}

	var out GetIngestsOutputType
	err = session.Request("GET", u.String(), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
