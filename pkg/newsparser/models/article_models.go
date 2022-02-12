package models

// Structure of the positions of an discovered article -- this reflects pretty much the API discovery type
// 	but has been defined separately since this package should be able to use anywhere else too
// 	furthermore the api might change (v1 -> v2) and this might imply all parsers implemented
// 	this also reduces imports for every parser
type DiscoveredArticle struct {
	ID          string
	URL         string
	ReleaseDate string
	Title       string
	Subtitle    string
	Description string
	Category    string
	Newspaper   string
}

type Newspaper struct {
	Name       string
	BaseURL    string
	LookupURLs []string
}

type Parser struct {
	// Name of the parser
	Name string
	// Version number of the parser
	Version string
	// Newspapers that can be parsed by this parser
	Newspapers []*Newspaper
}
