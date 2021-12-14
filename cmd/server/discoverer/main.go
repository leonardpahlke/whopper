package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/newsparser"
	"climatewhopper/pkg/whopperutil"
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// implementedDiscoveryServer is used to implement Discovery function.
type implementedDiscoveryServer struct {
	clients *whopperutil.WhopperServerClients
	config  *DiscovererServerConfig
	api.UnimplementedDiscovererServer
}

// DiscovererServerConfig configuration which will be injected as dapr sidecar information
type DiscovererServerConfig struct {
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
	config := DiscovererServerConfig{}
	whopperutil.SetViperCfg(fmt.Sprintf("dapr-%s-config", whopperutil.WhopperControllerDiscoverer), func() {}, &config)

	// create clients for server
	clients := whopperutil.GetWhopperClient(whopperutil.WhopperServerClientIn{
		LogLevel:      config.Spec.Template.Metadata.Annotations.LogLevel,
		SetDaprClient: true,
	})
	// close clients after server shutsdown
	defer clients.Close()

	// create grpc server (without starting it yet)
	serverInit, err := whopperutil.CreateGrpcServer(config.Spec.Template.Metadata.Annotations.AppPort)
	if err != nil {
		clients.Logger.Fatalw("could not create plain grpc server", "error", err)
	}

	// register grpc translator server
	api.RegisterDiscovererServer(serverInit.Server, &implementedDiscoveryServer{clients: &clients, config: &config})

	// start listening on port
	serverInit.StartListening(clients.Logger)
}

//
// Server Methods
//

func (s *implementedDiscoveryServer) Discover(ctx context.Context, in *api.DiscovererRequest) (*api.DiscovererResponse, error) {
	// run batch discovery
	unprocessedArticles, err := newsparser.BatchDiscovery(in.Info)
	if err != nil {
		s.clients.Logger.Errorw("could not run batch discovery")
		return nil, err
	}

	return &api.DiscovererResponse{
		Articles: unprocessedArticles,
		Head: &api.Head{
			Status:        api.Status_OK,
			StatusMessage: fmt.Sprintf("successfully discovered new articles, amount: %d", len(unprocessedArticles)),
			Timestamp:     timestamppb.Now(),
		},
	}, nil
}
