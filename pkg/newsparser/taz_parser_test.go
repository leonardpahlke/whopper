package newsparser

import (
	"climatewhopper/pkg/api"
	"testing"

	"github.com/foolin/pagser"
	"github.com/stretchr/testify/assert"
)

var tazDiscoveryData = []*api.InDiscovererInfo{{
	Newspaper: "taz",
	Url:       "https://taz.de/!t5204208/",
	LatestId:  0,
}, {
	Newspaper: "taz",
	Url:       "https://taz.de/!t5575293/",
	LatestId:  0,
}}

func TestDiscoverArticlesPos(t *testing.T) {
	p := pagser.New()
	for _, info := range tazDiscoveryData {
		articles, err := TazParser.DiscoverArticles(p, func() (string, error) {
			return GetWebsiteData(info.Url)
		}, 0)
		assert.NoError(t, err)
		assert.Greater(t, len(articles), 0)
	}
	// TODO: improve test
}

func TestBatchDiscoverArticlesPos(t *testing.T) {
	resp, err := BatchDiscovery(tazDiscoveryData)
	assert.NoError(t, err)
	assert.Greater(t, len(resp), 0)
	// TODO: improve test
}

func TestDiscoverArticlesNeg(t *testing.T) {
	// TODO: ...
}
