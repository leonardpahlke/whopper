package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/util"
	"climatewhopper/pkg/whopperutil"
	"context"
	"errors"
	"fmt"
	"net"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// implementedHubServer is used to implement RunHub function.
type implementedHubServer struct {
	logger     *zap.SugaredLogger
	daprClient dapr.Client
	api.UnimplementedHubServer
}

// RunHub implements the gRPC server function
// The hub is used as scheduler which kicks off processing
// 1. Start Discoverer
// 2. Start Argo Workflow (article processing)
func (s implementedHubServer) RunHub(ctx context.Context, in *api.HubRequest) (*api.HubResponse, error) {
	// TODO: implement RunHub
	return nil, errors.New("server not implemented yet")
}

func main() {
	logger := util.GetLogger(util.MatchLogLevel(util.WrapLogLevel(viper.GetString("LogLevel"))))

	// create dapr client
	client, err := dapr.NewClient()
	if err != nil {
		logger.Panicw("could not create dapr client", "error", err)
	}
	defer client.Close()

	// net listen
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("Port")))
	if err != nil {
		logger.Fatalw("failed to listen", "error", err)
	}

	// create new gRPC server
	s := grpc.NewServer()
	// Register server
	api.RegisterHubServer(s, &implementedHubServer{
		daprClient: client,
		logger:     util.GetLogger(zap.DebugLevel),
	})
	logger.Infow("server is listening", "address", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalw("failed to server", "error", err)
	}
}

func init() {
	util.SetViperCfg(string(whopperutil.WhopperControllerHub), func() {
		// set config defaults
		viper.SetDefault("Port", 50050)
		viper.SetDefault("DaprStoreName", "statestore")
		viper.SetDefault("LogLevel", util.Debug)
		// set flags
		pflag.Bool("test", false, "testmode")
	})
}
