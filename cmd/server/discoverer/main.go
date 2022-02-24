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

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	whopper "whopper/pkg/api/v1"
	"whopper/pkg/api/v1/discoverer"
	"whopper/pkg/newsparser"
	"whopper/pkg/util"
	"whopper/pkg/whopperutil"

	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"

	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// implementedDiscoveryServer is used to implement Discovery function.
type implementedDiscoveryServer struct {
	clients *whopperutil.WhopperServerClients
	config  *DiscovererServerConfig
	discoverer.UnimplementedDiscovererServer
}

// DiscovererServerConfig configuration which will be injected as dapr sidecar information
type DiscovererServerConfig struct {
	Port       int
	LogLevel   string
	StateStore string
}

const ServiceName = "Discoverer"

var (
	cfg                   = DiscovererServerConfig{}
	implementedNewspapers = newsparser.GetSupportedNewspapers()
	implementedParsers    = newsparser.GetAvailableParserIdentities()

	// supported*** represents collections that are available to look up
	// essentially these types are used from the endpoint GetNewspapers to return all valid newspapers without leaking information from the implemented versions
	SupportedNewspapers = []*whopper.Group{}
	SupportedParsers    = []*whopper.Parser{}
)

func init() {
	flag.IntVarP(&cfg.Port, "port", "p", 50051, "container port")
	flag.StringVar(&cfg.StateStore, "store", "statestore", "Statestore name that has been setup with dapr")
	flag.StringVar(&cfg.LogLevel, "log-level", string(util.Info), fmt.Sprintf("Log level of the container, %s", util.WrapLogLevels))
	flag.Parse()
	// filter supported Newspapers from implemented newspaper
	for _, newspaper := range implementedNewspapers {
		SupportedNewspapers = append(SupportedNewspapers, &whopper.Group{
			Name: newspaper.Name,
		})
	}
	// filter supported parsers from implemented parsers
	for _, parser := range implementedParsers {
		SupportedParsers = append(SupportedParsers, &whopper.Parser{
			Name: parser.Name,
		})
	}
}

func main() {
	// create clients for server
	clients := whopperutil.GetWhopperClient(whopperutil.WhopperServerClientIn{
		LogLevel:      cfg.LogLevel,
		SetDaprClient: true,
	})
	// close clients after server shutsdown
	defer clients.Close()

	clients.Logger.Debugw("configuration set to:", cfg)

	// create grpc server (without starting it yet)
	serverInit, err := whopperutil.CreateGrpcServer(strconv.Itoa(cfg.Port))
	if err != nil {
		clients.Logger.Fatalw("could not create plain grpc server", "error", err)
	}

	// register grpc discoverer server
	discoverer.RegisterDiscovererServer(serverInit.Server, &implementedDiscoveryServer{clients: &clients, config: &cfg})

	// start listening on port
	serverInit.StartListening(clients.Logger)
}

//
// API HANDLERS
//
// Discover articles of the specified newspaper with the selected parser
func (s *implementedDiscoveryServer) Discover(ctx context.Context, in *discoverer.DiscoverRequest) (*discoverer.DiscoverResponse, error) {
	// check parser & check if the newspaper is supported & check if newspaper is supported
	discovererArgs, err := verifyParserAndNewspaperInput(in.Newspaper, in.Parser)
	if err != nil {
		return &discoverer.DiscoverResponse{
			Articles: nil,
			Status: &whopper.Status{
				Code:    code.Code_INVALID_ARGUMENT,
				Message: "invalid input",
			},
		}, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	// discoveredArticles
	resp := []*discoverer.DiscoveredArticle{}
	for _, lookupPath := range discovererArgs.selectedNewspaper.LookupPaths {
		discoveredArticles, err := runDiscoverArticleTexts(s.clients.PagserClient, *discovererArgs.implParser, fmt.Sprintf("%s/%s", discovererArgs.selectedNewspaper.BaseURL, lookupPath))
		if err != nil {
			return &discoverer.DiscoverResponse{
				Articles: nil,
				Status: &whopper.Status{
					Code:    code.Code_CANCELLED,
					Message: "articles could not be discovered",
				},
			}, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
		}

		// run this function to store a discovered article text to the state store
		storedArticles, err := runStoreArticlesToStatestore(ctx, s.clients.DaprClient, discoveredArticles, in.Newspaper)
		if err != nil {
			marshalledData, err2 := json.Marshal(discoveredArticles)
			if err2 != nil {
				err = fmt.Errorf("err: %v and also could not marshal discoveredArticles err: %v", err, err2)
			}
			return &discoverer.DiscoverResponse{
				Articles: nil,
				Status: &whopper.Status{
					Code:    code.Code_CANCELLED,
					Message: "internal error storeing discovery results",
					Details: []*anypb.Any{{Value: marshalledData}},
				},
			}, status.Errorf(codes.Internal, "%v", err)
		}
		resp = append(resp, storedArticles...)
	}
	return &discoverer.DiscoverResponse{
		Articles: resp,
		Status: &whopper.Status{
			Code:    code.Code_OK,
			Message: "articles successfully discovered",
		},
	}, nil
}

// List api handler to return discovered articles
func (s *implementedDiscoveryServer) List(ctx context.Context, in *discoverer.ListDiscovererRequest) (*discoverer.ListDiscovererResponse, error) {
	// state, err := s.clients.DaprClient.QueryStateAlpha1(
	// 	ctx,
	// 	s.config.StateStore,
	// 	fmt.Sprintf("{'filter':{'EQ':{'value.state':'%s'}}}", in.Newspaper)
	// )
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// Request a discovered article from the database
func (s *implementedDiscoveryServer) Get(ctx context.Context, in *discoverer.GetDiscovererRequest) (*discoverer.DiscoveredArticle, error) {
	// request state from dapr state storage
	state, err := s.clients.DaprClient.GetState(ctx, s.config.StateStore, getKey(in.Id))
	if err != nil {
		return nil, errors.Wrap(err, "could not request state resource")
	}
	var article discoverer.DiscoveredArticle
	if err := proto.Unmarshal(state.Value, &article); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal state resource to object")
	}
	return &article, nil
}

// Return all newspapers that are valid to get requested
func (s *implementedDiscoveryServer) GetNewspapers(ctx context.Context, in *emptypb.Empty) (*discoverer.GetNewspapersResponse, error) {
	return &discoverer.GetNewspapersResponse{
		Newspapers: SupportedNewspapers,
	}, nil
}

// Return all newspapers that are valid to get requested
func (s *implementedDiscoveryServer) GetParsers(context.Context, *emptypb.Empty) (*discoverer.GetParsersResponse, error) {
	return &discoverer.GetParsersResponse{
		Parsers: SupportedParsers,
	}, nil
}

// Update a discovered article in the database
func (s *implementedDiscoveryServer) Update(ctx context.Context, in *discoverer.DiscoveredArticle) (*discoverer.DiscoveredArticle, error) {
	// request record which should get updated
	state, err := s.clients.DaprClient.GetState(ctx, s.config.StateStore, getKey(in.Id))
	if err != nil {
		return nil, errors.Wrap(err, "could not request state resource")
	}
	if state == nil {
		return nil, status.Errorf(codes.NotFound, "could not find state resource")
	}
	var storedArticle discoverer.DiscoveredArticle
	if err := proto.Unmarshal(state.Value, &storedArticle); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal state resource to object")
	}
	// copy meta information of the stored article before it can be overwritten to safe the created_at timestamp
	updateTime := timestamppb.Now()
	updatedStoredArticleMeta := whopper.Meta{
		CreatedAt: storedArticle.EntryMeta.CreatedAt,
		UpdatedAt: updateTime,
		Service:   ServiceName,
	}
	// overwrite the input meta
	in.EntryMeta = &updatedStoredArticleMeta
	// transform input message into bytes
	marshalledInput, err := proto.Marshal(in)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("could not marshal provided object to bytes, err: %v", err))
	}
	// overwrite retrieved object from the database with updated input fields (only updates if information is set)
	if err := proto.Unmarshal(marshalledInput, &storedArticle); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("could not unmarshal to stored resource, err: %v", err))
	}
	marshalledUpdatedArticle, err := proto.Marshal(&storedArticle)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("could not marshal updated object to bytes, err: %v", err))
	}
	// overwrite stored article with updated object
	if err := s.clients.DaprClient.SaveState(ctx, s.config.StateStore, getKey(in.Id), marshalledUpdatedArticle); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("could not store updated object, err: %v", err))
	}
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

// Delete a record on the statestore
func (s *implementedDiscoveryServer) Delete(ctx context.Context, in *discoverer.DeleteDiscoveredArticleRequest) (*emptypb.Empty, error) {
	if err := s.clients.DaprClient.DeleteState(ctx, s.config.StateStore, getKey(in.Id)); err != nil {
		return nil, status.Errorf(codes.NotFound, "could not find resource")
	}
	return &emptypb.Empty{}, nil
}

//
// HELPER FUNCTIONS
//

// this function defines database key formatting
func getKey(whopperID *whopper.ID) string {
	return fmt.Sprintf("disc-%s-%s", whopperID.Id, whopperID.Group.Name)
}
