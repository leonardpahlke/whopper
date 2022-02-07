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
	"strings"
	"whopper/pkg/api"

	"github.com/PuerkitoBio/goquery"
	"github.com/foolin/pagser"
	"github.com/pkg/errors"
)

//
// Discoverer Parser
//

// TazNewsParser taz parser struct
type TazNewsParser struct {
	newspaperName Newspaper
}

// TazParser is used to parse taz articles
var TazParser = TazNewsParser{
	newspaperName: "taz",
}

// GetName just returns the newspaper name which is used to match the parser
func (n TazNewsParser) GetName() Newspaper {
	return n.newspaperName
}

// DiscoverArticles used to scan taz newspaper category pages for articles
// if all articles should get returned, set stopAtID to 0 or any ID which does not match
func (n TazNewsParser) DiscoverArticles(p *pagser.Pagser, getWebsiteData func() (string, error), stopAtID int64) ([]*api.ArticleHead, error) {
	body, err := getWebsiteData()
	if err != nil {
		return nil, errors.Wrap(err, "could not get website data")
	}

	// parse website data into a struct which contains information about how to parse the website
	var data TazArticleDiscovery
	err = p.Parse(&data, body)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse taz article html data")
	}

	// transform taz specific data into a generic format
	articleHeads := []*api.ArticleHead{}
	for _, e := range data.Articles {
		if e.ID == stopAtID {
			// The website structures articles in order, this allows to break if the ID matches -- all articles after that have already been processed
			break
		}
		if e.ID != 0 {
			articleHeads = append(articleHeads, &api.ArticleHead{
				Id:          e.ID,
				Url:         e.URL,
				ReleaseDate: e.Date,
				Title:       e.Title,
				Subtitle:    e.SubTitle,
				Description: e.Description,
				Category:    data.Category,
			})
		}
	}

	return articleHeads, nil
}

// ParseArticle used to parse taz newspaper articles
func (n TazNewsParser) ParseArticle(p *pagser.Pagser, articleText *string) (*api.ArticleBody, error) {
	var data TazArticleTextParser
	err := p.Parse(&data, *articleText)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse taz article")
	}
	return &api.ArticleBody{
		OriginalText: strings.Join(data.Paragraphs, " "),
	}, nil
}

// TazArticleDiscovery represents the taz website overview
type TazArticleDiscovery struct {
	Category string                     `pagser:"title->CleanTazCategory()"`
	Articles []TazArticleDiscoveryEntry `pagser:"ul[role='directory'][class='news directory'] li"`
}

// TazArticleDiscoveryEntry represents one of the articles which is listed on the article overview page
type TazArticleDiscoveryEntry struct {
	ID          int64  `pagser:"meta[itemprop='cms-article-ID']->attr(content)"`
	URL         string `pagser:"a->AddTazURL()"`
	Date        string `pagser:"li[class='date']->RemoveSpaces()"`
	Title       string `pagser:"h3->text()"`
	SubTitle    string `pagser:"h4->text()"`
	Description string `pagser:"p->text()"`
}

// CleanTazCategory this method is used to clean up the taz category
func (d TazArticleDiscovery) CleanTazCategory(node *goquery.Selection, args ...string) (out interface{}, err error) {
	return strings.TrimSpace(strings.Replace(node.Text(), " - taz.de", "", 1)), nil
}

// AddTazURL this method is used to add the taz url to article reference path
func (d TazArticleDiscovery) AddTazURL(node *goquery.Selection, args ...string) (out interface{}, err error) {
	return fmt.Sprintf("https://taz.de%s", node.AttrOr("href", "")), nil
}

// RemoveSpaces this method is used to remove all spaces from a text
func (d TazArticleDiscovery) RemoveSpaces(node *goquery.Selection, args ...string) (out interface{}, err error) {
	return strings.ReplaceAll(node.Text(), " ", ""), nil
}

//
// Article Parser
///

// TazArticleTextParser represents the structure of a taz article
type TazArticleTextParser struct {
	Paragraphs []string `pagser:"article[class='sectbody'][itemprop='articleBody'] p[xmlns='']->text()"`
}
