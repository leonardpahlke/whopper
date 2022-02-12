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

// TODO: write unit tests
// var tazDiscoveryData = []*api.InDiscovererInfo{{
// 	Newspaper: "taz",
// 	Url:       "https://taz.de/!t5204208/",
// 	LatestId:  0,
// }, {
// 	Newspaper: "taz",
// 	Url:       "https://taz.de/!t5575293/",
// 	LatestId:  0,
// }}

// func TestDiscoverArticlesPos(t *testing.T) {
// 	p := pagser.New()
// 	for _, info := range tazDiscoveryData {
// 		articles, err := TazParser.DiscoverArticles(p, func() (string, error) {
// 			return GetWebsiteData(info.Url)
// 		}, 0)
// 		assert.NoError(t, err)
// 		assert.Greater(t, len(articles), 0)
// 	}
// 	// TODO: improve test
// }

// func TestBatchDiscoverArticlesPos(t *testing.T) {
// 	resp, err := BatchDiscovery(tazDiscoveryData)
// 	assert.NoError(t, err)
// 	assert.Greater(t, len(resp), 0)
// 	// TODO: improve test
// }

// func TestDiscoverArticlesNeg(t *testing.T) {
// 	// TODO: ...
// }
