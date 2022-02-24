package models

import "encoding/json"

// Structure of the positions of an discovered article -- this reflects pretty much the API discovery type
// 	but has been defined separately since this package should be able to use anywhere else too
// 	furthermore the api might change (v1 -> v2) and this might imply all parsers implemented
// 	this also reduces imports for every parser

//-------------------------
type DiscoveredArticle struct {
	ID          string `json:"id,omitempty"`
	URL         string `json:"url,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
	Title       string `json:"title,omitempty"`
	Subtitle    string `json:"subtitle,omitempty"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
	Newspaper   string `json:"newspaper,omitempty"`
}

func UnmarshalDiscoveredArticle(data []byte) (DiscoveredArticle, error) {
	var r DiscoveredArticle
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscoveredArticle) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//-------------------------
type Newspaper struct {
	Name        string   `json:"name,omitempty"`
	BaseURL     string   `json:"base_url,omitempty"`
	LookupPaths []string `json:"lookup_urls,omitempty"`
}

func UnmarshalNewspaper(data []byte) (Newspaper, error) {
	var r Newspaper
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Newspaper) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//-------------------------
type Parser struct {
	// Name of the parser
	Name string `json:"name,omitempty"`
	// Version number of the parser
	Version string `json:"version,omitempty"`
	// Newspapers that can be parsed by this parser
	Newspapers []*Newspaper `json:"newspapers,omitempty"`
}

func UnmarshalParser(data []byte) (Parser, error) {
	var r Parser
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Parser) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
