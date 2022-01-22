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
	"climatewhopper/pkg/api"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/foolin/pagser"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// NewsParser is used to define a newspaper parser
type NewsParser interface {
	GetName() Newspaper
	DiscoverArticles(p *pagser.Pagser, getWebsiteData func() (string, error), stopAtID int64) ([]*api.ArticleHead, error)
	ParseArticle(p *pagser.Pagser, articleText *string) (*api.ArticleBody, error)
}

// Newspaper is used to match a parser to a newspaper
type Newspaper string

// NewsParsers contains all the parsers that are available
var NewsParsers = []NewsParser{TazParser}

// DiscoveryLookup this type is used if multiple discoveries should happen
type DiscoveryLookup struct {
	Newspaper   Newspaper
	Error       error
	ArticleHead []*api.ArticleHead
}

// GetNewspaperParser is used to simplify how to get the right parser for a newspaper article
func GetNewspaperParser(newspaper Newspaper) (NewsParser, error) {
	for _, parser := range NewsParsers {
		if parser.GetName() == newspaper {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("could not find a parser for newspaper %s", newspaper)
}

// BatchDiscovery used to run multiple concurrent website discoveries
func BatchDiscovery(discoveryBatch []*api.InDiscovererInfo) ([]*api.ArticleHead, error) {
	p := pagser.New()
	discoverBatch := func(done <-chan interface{}, in ...*api.InDiscovererInfo) <-chan DiscoveryLookup {
		chanLookups := make(chan DiscoveryLookup)
		go func() {
			defer close(chanLookups)
			for i := range in {
				// get article parser by newspaper and run DiscoverArticles
				discoveredArticleHeads := []*api.ArticleHead{}
				parser, err := GetNewspaperParser(Newspaper(in[i].Newspaper))
				// if no newspaper parser could be found the DiscoverArticle method will not get executed (skipped as it will error anyways)
				if err == nil {
					discoveredArticleHeads, err = parser.DiscoverArticles(p, func() (string, error) {
						return GetWebsiteData(in[i].Url)
					}, in[i].LatestId)
				}
				select {
				case <-done:
					return
				case chanLookups <- DiscoveryLookup{
					Newspaper:   Newspaper(in[i].Newspaper),
					Error:       err,
					ArticleHead: discoveredArticleHeads,
				}:
				}
			}
		}()
		return chanLookups
	}

	done := make(chan interface{})
	defer close(done)
	discoveredArticles := []*api.ArticleHead{}
	var err error

	// Collect data from buffered channel
	for lookups := range discoverBatch(done, discoveryBatch...) {
		if lookups.Error != nil {
			err = multierror.Append(err, errors.Wrapf(lookups.Error, "error discovering articles for newspaper %s", lookups.Newspaper))
		} else {
			discoveredArticles = append(discoveredArticles, lookups.ArticleHead...)
		}
	}
	return discoveredArticles, err
}

// GetWebsiteData simple http get wrap
func GetWebsiteData(webURL string) (string, error) {
	resp, err := http.Get(strings.TrimSpace(webURL))
	if err != nil {
		return "", errors.Wrap(err, "could not perform http request with http.Get")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "could not read http response body with ioutil.ReadAll")
	}
	return string(body), nil
}
