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

package taz

import (
	"fmt"
	"strconv"
	"strings"
	"whopper/pkg/newsparser/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/foolin/pagser"
	"github.com/pkg/errors"
)

//
// TAZ NEWSPAPER PARSER
//

var Newspaper = models.Newspaper{
	Name:       "taz",
	BaseURL:    "https://taz.de",
	LookupURLs: []string{}, // TODO: not sure about this one
}

// NewsParserV1 taz parser struct
type NewsParserV1 struct{}

// GetName just returns the newspaper name which is used to match the parser
func (n NewsParserV1) GetIdentity() models.Parser {
	return models.Parser{
		Name:       fmt.Sprintf("%s-parser", Newspaper.Name),
		Version:    "1.0",
		Newspapers: []*models.Newspaper{&Newspaper},
	}
}

// DiscoverArticles used to scan taz newspaper category pages for articles
// if all articles should get returned, set stopAtID to 0 or any ID which does not match
func (n NewsParserV1) DiscoverArticles(p *pagser.Pagser, getWebsiteData func() (string, error), stopAtID int64) ([]*models.DiscoveredArticle, error) {
	body, err := getWebsiteData()
	if err != nil {
		return nil, errors.Wrap(err, "could not get website data")
	}

	// parse website data into a struct which contains information about how to parse the website
	var data ArticleDiscovery
	err = p.Parse(&data, body)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse taz article html data")
	}

	// transform taz specific data into a generic format
	articleHeads := []*models.DiscoveredArticle{}
	for _, e := range data.Articles {
		if e.ID == stopAtID {
			// The website structures articles in order, this allows to break if the ID matches -- all articles after that have already been processed
			break
		}
		if e.ID != 0 {
			articleHeads = append(articleHeads, &models.DiscoveredArticle{
				ID:          strconv.FormatInt(e.ID, 10),
				URL:         e.URL,
				ReleaseDate: e.Date,
				Title:       e.Title,
				Subtitle:    e.SubTitle,
				Description: e.Description,
				Category:    data.Category,
				Newspaper:   "taz",
			})
		}
	}

	return articleHeads, nil
}

// ParseArticle used to parse taz newspaper articles
func (n NewsParserV1) ParseArticle(p *pagser.Pagser, articleText *string) (*string, error) {
	var data ArticleTextParser
	err := p.Parse(&data, *articleText)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse taz article")
	}
	originalText := strings.Join(data.Paragraphs, " ")
	return &originalText, nil
}

// ArticleDiscovery represents the taz website overview
type ArticleDiscovery struct {
	Category string                  `pagser:"title->CleanTazCategory()"`
	Articles []ArticleDiscoveryEntry `pagser:"ul[role='directory'][class='news directory'] li"`
}

// ArticleDiscoveryEntry represents one of the articles which is listed on the article overview page
type ArticleDiscoveryEntry struct {
	ID          int64  `pagser:"meta[itemprop='cms-article-ID']->attr(content)"`
	URL         string `pagser:"a->AddTazURL()"`
	Date        string `pagser:"li[class='date']->RemoveSpaces()"`
	Title       string `pagser:"h3->text()"`
	SubTitle    string `pagser:"h4->text()"`
	Description string `pagser:"p->text()"`
}

// CleanCategory this method is used to clean up the taz category
func (d ArticleDiscovery) CleanCategory(node *goquery.Selection, args ...string) (out interface{}, err error) {
	return strings.TrimSpace(strings.Replace(node.Text(), " - taz.de", "", 1)), nil
}

// AddURL this method is used to add the taz url to article reference path
func (d ArticleDiscovery) AddURL(node *goquery.Selection, args ...string) (out interface{}, err error) {
	return fmt.Sprintf("https://taz.de%s", node.AttrOr("href", "")), nil
}

// RemoveSpaces this method is used to remove all spaces from a text
func (d ArticleDiscovery) RemoveSpaces(node *goquery.Selection, args ...string) (out interface{}, err error) {
	return strings.ReplaceAll(node.Text(), " ", ""), nil
}

//
// Article Parser
///

// ArticleTextParser represents the structure of a taz article
type ArticleTextParser struct {
	Paragraphs []string `pagser:"article[class='sectbody'][itemprop='articleBody'] p[xmlns='']->text()"`
}
