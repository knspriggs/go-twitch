package twitch

// IngestType -
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

// GetIngestsOutputType -
type GetIngestsOutputType struct {
	Ingests []IngestType      `json:"ingests"`
	Links   map[string]string `json:"_links"`
}

// GetIngests -
func (session *Session) GetIngests() (*GetIngestsOutputType, error) {
	var out GetIngestsOutputType
	err := session.Request("GET", "/ingests", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
