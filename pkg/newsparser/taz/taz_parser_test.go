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
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"whopper/pkg/newsparser/models"

	"github.com/foolin/pagser"
	"github.com/stretchr/testify/assert"
)

// TODO: write unit tests
// var tazDiscoveryData = []*whopper.InDiscovererInfo{{
// 	Newspaper: "taz",
// 	Url:       "https://taz.de/!t5204208/",
// 	LatestId:  0,
// }, {
// 	Newspaper: "taz",
// 	Url:       "https://taz.de/!t5575293/",
// 	LatestId:  0,
// }}

// var tazParser = NewsParserV1{}

// 5204208-disco.txt
// 5575293-disco.txt
// gespraech-mit-der-zivilgesellschaft-info.txt
// klimabuergerinnenrat-startet-info.txt

type data struct {
	sourceFilePath string
	targetFilePath string
	latestID       string
}

const testFolderName = "test-data"

var discoveryTestdataPos = []data{
	{
		sourceFilePath: "5204208-0-disco.s.txt",
		targetFilePath: "5204208-0-disco.t.txt",
		latestID:       "-",
	},
	{
		sourceFilePath: "5204208-0-disco.s.txt",
		targetFilePath: "5204208-5787788-disco.t.txt",
		latestID:       "5787788",
	},
	{
		sourceFilePath: "5575293-0-disco.s.txt",
		targetFilePath: "5575293-0-disco.t.txt",
		latestID:       "-",
	},
	{
		sourceFilePath: "5575293-0-disco.s.txt",
		targetFilePath: "5575293-5757341-disco.t.txt",
		latestID:       "5757341",
	},
}

func TestDiscoverArticlesPos(t *testing.T) {
	p := pagser.New()
	newsParser := NewsParserV1{}

	for _, data := range discoveryTestdataPos {
		sourceFileBytes, err := readFile(testFolderName + "/" + data.sourceFilePath)
		assert.NoError(t, err, "could not read testfile")
		discoveredArticles, err := newsParser.DiscoverArticles(p, &sourceFileBytes, data.latestID)
		assert.NoError(t, err, "could not discover articles")
		marshalledArticles, err := json.Marshal(discoveredArticles)
		assert.NoError(t, err, "could not transform discovered article data into json format")
		targetFileBytes, err := readFile(testFolderName + "/" + data.targetFilePath)
		assert.NoError(t, err, "could not read testfile")
		assert.Equal(t, targetFileBytes, marshalledArticles, "discovered article data is not the expected result")

		// NOTE: This can be used to generate the output test files
		// err = ioutil.WriteFile(testFolderName+"/"+data.targetFilePath, marshalledArticles, 0644)
		// assert.NoError(t, err)
	}
}

var discoveryTestdataNegNoErr = []data{
	{
		sourceFilePath: "neg-sueddeutsche-disco.txt",
		targetFilePath: "-",
		latestID:       "-",
	},
}

func TestDiscoverArticlesNegNoErr(t *testing.T) {
	p := pagser.New()
	newsParser := NewsParserV1{}

	for _, data := range discoveryTestdataNegNoErr {
		sourceFileBytes, err := readFile(testFolderName + "/" + data.sourceFilePath)
		assert.NoError(t, err, "could not read testfile")
		discoveredArticles, err := newsParser.DiscoverArticles(p, &sourceFileBytes, data.latestID)
		assert.NoError(t, err, "faulty input which should not raise an error")
		assert.Equal(t, discoveredArticles, []*models.DiscoveredArticle{}, "no articles should get discovered with faulty input")
	}
}

var parseTestdataPos = []data{
	{
		sourceFilePath: "in-2f3653d37ddd02631792c47b.txt",
		targetFilePath: "out-2f3653d37ddd02631792c47b.txt",
		latestID:       "-",
	},
	{
		sourceFilePath: "in-14d155d61a2519c66194953hf.txt",
		targetFilePath: "out-14d155d61a2519c66194953hf.txt",
		latestID:       "-",
	},
}

func TestParseArticlesPos(t *testing.T) {
	p := pagser.New()
	newsParser := NewsParserV1{}

	for _, data := range parseTestdataPos {
		sourceFileBytes, err := readFile(testFolderName + "/" + data.sourceFilePath)
		sourceFileString := string(sourceFileBytes)
		assert.NoError(t, err, "could not read testfile")
		parsedArticle, err := newsParser.ParseArticle(p, &sourceFileString)
		assert.NoError(t, err, "could not discover articles")
		marshalledArticle, err := json.Marshal(parsedArticle)
		assert.NoError(t, err, "could not transform discovered article data into json format")
		targetFileBytes, err := readFile(testFolderName + "/" + data.targetFilePath)
		assert.NoError(t, err, "could not read testfile")
		assert.Equal(t, targetFileBytes, marshalledArticle, "discovered article data is not the expected result")

		// NOTE: This can be used to generate the output test files
		// err = ioutil.WriteFile(testFolderName+"/"+data.targetFilePath, marshalledArticles, 0644)
		// assert.NoError(t, err)
	}
}

var parseTestdataNegNoErr = []data{
	{
		sourceFilePath: "neg-sueddeutsche-disco.txt",
		targetFilePath: "-",
		latestID:       "-",
	},
}

func TestParseArticlesNegNoErr(t *testing.T) {
	p := pagser.New()
	newsParser := NewsParserV1{}

	for _, data := range parseTestdataNegNoErr {
		sourceFileBytes, err := readFile(testFolderName + "/" + data.sourceFilePath)
		sourceFileString := string(sourceFileBytes)
		assert.NoError(t, err, "could not read testfile")
		parsedArticle, err := newsParser.ParseArticle(p, &sourceFileString)
		assert.NoError(t, err, "faulty input which should not raise an error")
		assert.Equal(t, "", *parsedArticle, "no articles should get discovered with faulty input")
	}
}

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return ioutil.ReadAll(file)
}
