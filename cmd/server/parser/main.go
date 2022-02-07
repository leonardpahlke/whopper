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
	"fmt"
	"whopper/pkg/api"
	"whopper/pkg/newsparser"
	"whopper/pkg/whopperutil"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// implementedParserServer is used to implement Parse function.
type implementedParserServer struct {
	clients *whopperutil.WhopperServerClients
	config  *ParserServerConfig
	api.UnimplementedParserServer
}

// ParserServerConfig configuration which will be injected as dapr sidecar information
type ParserServerConfig struct {
	Spec struct {
		Template struct {
			Metadata struct {
				Annotations struct {
					AppID       string `mapstructure:"dapr.io/app-id"`
					AppProtocol string `mapstructure:"dapr.io/app-protocol"`
					AppPort     string `mapstructure:"dapr.io/app-port"`
					LogLevel    string `mapstructure:"dapr.io/app-level"`
				} `mapstructure:"annotations"`
			} `mapstructure:"metadata"`
		} `mapstructure:"template"`
	} `mapstructure:"spec"`
}

func main() {
	// parse configuration file with viper
	config := ParserServerConfig{}
	whopperutil.SetViperCfg(fmt.Sprintf("dapr-%s-config", whopperutil.WhopperEngineParser), func() {}, &config)

	// create clients for server
	clients := whopperutil.GetWhopperClient(whopperutil.WhopperServerClientIn{
		LogLevel:        config.Spec.Template.Metadata.Annotations.LogLevel,
		SetDaprClient:   true,
		SetPagserClient: true,
	})
	// close clients after server shutsdown
	defer clients.Close()

	// create grpc server (without starting it yet)
	serverInit, err := whopperutil.CreateGrpcServer(config.Spec.Template.Metadata.Annotations.AppPort)
	if err != nil {
		clients.Logger.Fatalw("could not create plain grpc server", "error", err)
	}

	// register grpc translator server
	api.RegisterParserServer(serverInit.Server, &implementedParserServer{clients: &clients, config: &config})

	// start listening on port
	serverInit.StartListening(clients.Logger)
}

//
// Server Methods
//

// Parse function extends gRPC parser server function and is used to parse a newspaper article text
func (s implementedParserServer) Parse(ctx context.Context, in *api.ParserRequest) (*api.ParserResponse, error) {
	s.clients.Logger.Info("run parser")
	parser, err := newsparser.GetNewspaperParser(newsparser.Newspaper(in.Newspaper))
	if err != nil {
		return &api.ParserResponse{
			Id:        in.Id,
			Newspaper: in.Newspaper,
			Head:      &api.Head{Status: api.Status_ERROR, StatusMessage: "could not find a parser to parse raw website data", Timestamp: timestamppb.Now()},
		}, err
	}
	articleBody, err := parser.ParseArticle(s.clients.PagserClient, &in.RawArticle)
	if err != nil {
		return &api.ParserResponse{
			Id:        in.Id,
			Newspaper: in.Newspaper,
			Head:      &api.Head{Status: api.Status_ERROR, StatusMessage: "could not parse raw website data", Timestamp: timestamppb.Now()},
		}, err
	}
	return &api.ParserResponse{
		Id:        in.Id,
		Newspaper: in.Newspaper,
		Text:      articleBody,
		Head:      &api.Head{Status: api.Status_OK, StatusMessage: "parsed article text from raw article data", Timestamp: timestamppb.Now()},
	}, nil
}
