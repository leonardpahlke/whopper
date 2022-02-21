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

package whopperutil

import (
	"context"
	"fmt"
	"net"
	"whopper/pkg/newsparser/models"
	"whopper/pkg/util"

	whopper "whopper/pkg/api/v1"

	language "cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/translate"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/foolin/pagser"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Package with whopper related utility

// WhopperServerName type used for gRPC servers of this project
type WhopperServerName string

var (
	// WhopperControllerDiscoverer discoverer server name
	WhopperControllerDiscoverer WhopperServerName = "discoverer"
	// WhopperControllerHub hub server name
	WhopperControllerHub WhopperServerName = "hub"
	// WhopperEngineParser parser server name
	WhopperEngineParser WhopperServerName = "parser"
	// WhopperEngineTranslator translator server name
	WhopperEngineTranslator WhopperServerName = "translator"
	// WhopperEngineAnalyzer analyzer server name
	WhopperEngineAnalyzer WhopperServerName = "analyzer"

	// WhopperServers available whopper servers
	WhopperServers = []WhopperServerName{WhopperControllerDiscoverer, WhopperControllerHub, WhopperEngineParser, WhopperEngineTranslator, WhopperEngineAnalyzer}
)

// WhopperServerClientIn define how to build the WhopperServerMeta with BuildWhopperServer
type WhopperServerClientIn struct {
	LogLevel              string
	SetDaprClient         bool
	SetGcpNlpClient       bool
	SetPagserClient       bool
	SetGcpTranslateClient bool
}

// WhopperServerClients meta information like logger, clients, ... which can be set using BuildWhopperServer
type WhopperServerClients struct {
	Ctx                context.Context
	Logger             *zap.SugaredLogger
	DaprClient         dapr.Client
	GcpLanguageClient  *language.Client
	GcpTranslateClient *translate.Client
	PagserClient       *pagser.Pagser
}

// GetWhopperClient used to simplify how to setup a whopper server
func GetWhopperClient(t WhopperServerClientIn) WhopperServerClients {
	meta := WhopperServerClients{Ctx: context.Background()}

	// create a new logger client
	meta.Logger = util.GetLogger(util.MatchLogLevel(util.WrapLogLevel(t.LogLevel)))

	// Create a dapr client
	if t.SetDaprClient {
		client, err := dapr.NewClient()
		if err != nil {
			meta.Logger.Fatalw("could not create dapr client", "error", err)
		}
		meta.DaprClient = client
	}

	// Creates a gcp ML Client
	if t.SetGcpNlpClient {
		gcpNlpClient, err := language.NewClient(meta.Ctx)
		if err != nil {
			meta.Logger.Fatalw("could not create a gcp language client", "error", err)
		}
		meta.GcpLanguageClient = gcpNlpClient
	}

	// Create a gcp translate Client
	if t.SetGcpTranslateClient {
		translateClient, err := translate.NewClient(meta.Ctx)
		if err != nil {
			meta.Logger.Fatalw("could not create a gcp translate client", "error", err)
		}
		meta.GcpTranslateClient = translateClient
	}

	// Create a pagser client
	if t.SetPagserClient {
		meta.PagserClient = pagser.New()
	}

	meta.Logger.Debug("clients have been setup")

	return meta
}

// WhopperServerInit used to simplify how to setup a grpc server
type WhopperServerInit struct {
	Server   *grpc.Server
	listener net.Listener
}

// CreateGrpcServer used to setup a grpc server on
func CreateGrpcServer(grpcPort string) (*WhopperServerInit, error) {
	// net listen
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}

	// create new gRPC server
	server := grpc.NewServer()
	return &WhopperServerInit{
		Server:   server,
		listener: lis,
	}, nil
}

// StartListening is used to start the grpc server
func (w *WhopperServerInit) StartListening(logger *zap.SugaredLogger) {
	logger.Infow("server is listening", "address", w.listener.Addr())
	if err := w.Server.Serve(w.listener); err != nil {
		logger.Fatalw("failed to server", "error", err)
	}
}

// Close this method should be called with defer to close created clients
func (m *WhopperServerClients) Close() {
	// close call clients that have been created
	if m.DaprClient != nil {
		defer m.DaprClient.Close()
	}
	if m.GcpLanguageClient != nil {
		defer m.GcpLanguageClient.Close()
	}
	if m.GcpTranslateClient != nil {
		defer m.GcpTranslateClient.Close()
	}
}

// DefaultConfigFilenameSuffix name of the configuration file used to configure whopper servers
const DefaultConfigFilenameSuffix = "-config"

// SetViperCfg simplify how to setup the viper configuration
// func SetViperCfg(configName string, setViperDefaults func(), config interface{}) {
// 	// set config meta
// 	viper.SetConfigName(configName)
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath("./climate-whopper/configs")
// 	viper.AddConfigPath("$HOME/climate-whopper/configs")
// 	viper.AddConfigPath("./configs")
// 	// set config defaults
// 	setViperDefaults()
// 	// bind flags
// 	pflag.Parse()
// 	err := viper.BindPFlags(pflag.CommandLine)
// 	if err != nil {
// 		panic(fmt.Errorf("fatal error binding flags: %w", err))
// 	}
// 	// read config
// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		panic(fmt.Errorf("fatal error config file: %w", err))
// 	}

// 	if err := viper.Unmarshal(config); err != nil {
// 		panic(fmt.Errorf("unable to decode into struct: %w", err))
// 	}
// }

// TranslateNewspaperDefinitions is used to convert between the API and Parser.Model types
// 	the reasoning behind this is briefly descirbed in the models file under the newsparser dir
func TranslateNewspaperDefinitions(newspaper []*models.Newspaper) []*whopper.Group {
	w := []*whopper.Group{}
	for _, e := range newspaper {
		w = append(w, &whopper.Group{
			Name: e.Name,
		})
	}
	return w
}

// TranslateParserDefinitions is used to convert between the API and Parser.Model types
func TranslateParserDefinitions(parsers []*models.Parser) []*whopper.Parser {
	p := []*whopper.Parser{}
	for _, e := range parsers {
		p = append(p, &whopper.Parser{
			Name: e.Name,
		})
	}
	return p
}
