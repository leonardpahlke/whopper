package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/whopperutil"
	"context"
	"errors"
	"fmt"
)

// implementedHubServer is used to implement RunHub function.
type implementedHubServer struct {
	clients *whopperutil.WhopperServerClients
	config  *HubServerConfig
	api.UnimplementedHubServer
}

// HubServerConfig configuration which will be injected as dapr sidecar information
type HubServerConfig struct {
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
	config := HubServerConfig{}
	whopperutil.SetViperCfg(fmt.Sprintf("dapr-%s-config", whopperutil.WhopperControllerHub), func() {}, &config)

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
	api.RegisterHubServer(serverInit.Server, &implementedHubServer{clients: &clients, config: &config})

	// start listening on port
	serverInit.StartListening(clients.Logger)
}

//
// Server Methods
//

// RunHub implements the gRPC server function
// The hub is used as scheduler which kicks off processing
// 1. Start Discoverer
// 2. Start Argo Workflow (article processing)
func (s implementedHubServer) RunHub(ctx context.Context, in *api.HubRequest) (*api.HubResponse, error) {
	// TODO: implement RunHub
	return nil, errors.New("not implemented yet")
}
