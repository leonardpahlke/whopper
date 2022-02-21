package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
	whopper "whopper/pkg/api/v1"
	"whopper/pkg/api/v1/discoverer"
	"whopper/pkg/newsparser"
	"whopper/pkg/newsparser/models"

	dapr "github.com/dapr/go-sdk/client"

	"github.com/foolin/pagser"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type discovererArgs struct {
	selectedNewspaper *models.Newspaper
	selectedParser    *models.Parser
	implParser        *newsparser.INewsParser
}

// This function is used to check if the provided parsing information is valid
func verifyParserAndNewspaperInput(newspaper *whopper.Group, parser *whopper.Parser) (*discovererArgs, error) {
	if newspaper.Name != "" {
		return nil, fmt.Errorf("newspaper is empty")
	}
	if parser.Name != "" {
		return nil, fmt.Errorf("parser is empty")
	}

	var (
		parserAndNewspaperValid = false
		newspaperToUse          *models.Newspaper
		implementedParserToUse  *newsparser.INewsParser
		parserDefinitionToUse   *models.Parser
	)
	// check if the parser is implemented and if the implemented parser supports the specified newspaper
	for _, implementedParser := range implementedParsers {
		if parser.Name != implementedParser.Name {
			continue
		}
		parserAndNewspaperValid = true
		for _, implementedNewspaper := range implementedNewspapers {
			if newspaper.Name == implementedNewspaper.Name {
				newspaperToUse = implementedNewspaper
				break
			}
		}
		if newspaperToUse == nil {
			return nil, fmt.Errorf("specified newspaper is not supported")
		}
		parserDefinitionToUse = implementedParser
		break
	}
	if !parserAndNewspaperValid {
		return nil, fmt.Errorf("specified parser is not supported")
	}
	return &discovererArgs{
		selectedNewspaper: newspaperToUse,
		selectedParser:    parserDefinitionToUse,
		implParser:        implementedParserToUse,
	}, nil
}

// discover articles from a newspaper URL
func runDiscoverArticleTexts(p *pagser.Pagser, parser newsparser.INewsParser, url string) ([]*models.DiscoveredArticle, error) {
	// send http request to URL to retrieve article information
	resp, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not request website err: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body err: %v", err)
	}
	// This should be looked up dynamic
	latestID := "0"
	discoveredArticles, err := parser.DiscoverArticles(p, &body, latestID)
	if err != nil {
		return nil, fmt.Errorf("could not discover articles: %v", err)
	}
	return discoveredArticles, nil
}

// used to store bare parsed articles to the state store
func runStoreArticlesToStatestore(ctx context.Context, daprClient dapr.Client, discoveredArticles []*models.DiscoveredArticle, newspaper *whopper.Group) ([]*discoverer.DiscoveredArticle, error) {
	storedArticles := []*discoverer.DiscoveredArticle{}
	for _, discoveredArticle := range discoveredArticles {
		// convert response to whooper format
		transformedDiscoveredArticle := discoverer.DiscoveredArticle{
			Id: &whopper.ID{
				Id: discoveredArticle.ID,
				Group: &whopper.Group{
					Name: newspaper.Name,
				},
			},
			Url:         discoveredArticle.URL,
			ReleaseDate: discoveredArticle.ReleaseDate,
			Title:       discoveredArticle.Title,
			Subtitle:    discoveredArticle.Subtitle,
			Description: discoveredArticle.Description,
			Category:    discoveredArticle.Category,
			Newspaper: &whopper.Group{
				Name: newspaper.Name,
			},
			EntryMeta: &whopper.Meta{
				CreatedAt: timestamppb.New(time.Now()),
				UpdatedAt: timestamppb.New(time.Now()),
				Service:   ServiceName,
			},
		}

		protoArticle, err := proto.Marshal(&transformedDiscoveredArticle)
		if err != nil {
			return nil, fmt.Errorf("could not transform data to protobuf message: %v", err)
		}
		// store article in state store
		err = daprClient.SaveState(
			ctx,
			cfg.StateStore,
			getKey(&whopper.ID{
				Id: discoveredArticle.ID,
				Group: &whopper.Group{
					Name: newspaper.Name,
				},
			}),
			protoArticle,
		)
		if err != nil {
			return nil, fmt.Errorf("could store state: %v", err)
		}
		// add article to response list
		storedArticles = append(storedArticles, &transformedDiscoveredArticle)
	}
	return storedArticles, nil
}
