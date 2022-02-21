// Copyright 2022 Leonard Vincent Simon Pahlke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package newsparser

import (
	"fmt"
	"whopper/pkg/newsparser/models"
	"whopper/pkg/newsparser/taz"

	"github.com/foolin/pagser"
)

// INewsParser is used to define a newspaper parser
type INewsParser interface {
	GetIdentity() models.Parser
	DiscoverArticles(p *pagser.Pagser, websiteRaw *[]byte, omitResultsAfterID string) ([]*models.DiscoveredArticle, error)
	ParseArticle(p *pagser.Pagser, articleText *string) (*string, error)
}

// ImplementedNewsParsers contains all the parsers that are available
var ImplementedNewsParsers = []INewsParser{taz.NewsParserV1{}}

// GetAvailableParserIdentities is used every implemeneted parser identification
func GetAvailableParserIdentities() []*models.Parser {
	parserIDs := []*models.Parser{}
	for _, e := range ImplementedNewsParsers {
		identity := e.GetIdentity()
		parserIDs = append(parserIDs, &identity)
	}
	return parserIDs
}

// GetAvailableParserIdentities is used every implemeneted parser identification
func GetSupportedNewspapers() []*models.Newspaper {
	// loop over all parsers and sum up all the supported newspapers
	// 	a map is used to avoid duplicates
	supportedNewspapersMap := map[string]*models.Newspaper{}
	for _, e := range ImplementedNewsParsers {
		for _, newspaper := range e.GetIdentity().Newspapers {
			supportedNewspapersMap[newspaper.Name] = newspaper
		}
	}
	supportedNewspapers := make([]*models.Newspaper, len(supportedNewspapersMap))
	i := 0
	for _, v := range supportedNewspapersMap {
		supportedNewspapers[i] = v
		i++
	}
	return supportedNewspapers
}

// GetNewspaperParser is used to simplify how to get the right parser for a newspaper article
func GetNewspaperParser(parserName, parserVersion string) (INewsParser, error) {
	for _, parser := range ImplementedNewsParsers {
		parserID := parser.GetIdentity()
		if parserID.Name == parserName && parserID.Version == parserVersion {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("could not find a parser implementation, %v", parserName)
}
