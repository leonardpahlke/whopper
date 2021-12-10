package newsparser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/foolin/pagser"
	"github.com/pkg/errors"
)

//
// Discoverer Parser
//

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

// TODO: define generic Discoverer interface which gets extended by TazDiscoverer (and others in future)

// TazDiscoverer is used to parse a category page on taz.de (like https://taz.de/!t5204208/)
func TazDiscoverer(p *pagser.Pagser, url string) (*TazArticleDiscovery, error) {
	// TODO: refactor this after another parser has been added
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "could not perform http request with http.Get")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read http response body with ioutil.ReadAll")
	}

	var data TazArticleDiscovery
	err = p.Parse(&data, string(body))
	if err != nil {
		return nil, errors.Wrap(err, "could not parse taz article html data")
	}

	unparsedArticles := TazArticleDiscovery{}
	for _, e := range data.Articles {
		if e.ID != 0 {
			unparsedArticles.Articles = append(unparsedArticles.Articles, e)
		}
	}

	return &unparsedArticles, nil
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

// TODO: ... parse taz article text see api.ArticleBody
