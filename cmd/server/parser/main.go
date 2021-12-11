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

// implementedParserServer is used to implement Parse function.
type implementedParserServer struct {
	logger     *zap.SugaredLogger
	daprClient dapr.Client
	api.UnimplementedParserServer
}

// Parse function extends gRPC parser server function and is used to parse a newspaper article text
func (s implementedParserServer) Parse(context.Context, *api.ParserRequest) (*api.ParserResponse, error) {
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
	api.RegisterParserServer(s, &implementedParserServer{
		daprClient: client,
		logger:     util.GetLogger(zap.DebugLevel),
	})
	logger.Infow("server is listening", "address", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalw("failed to server", "error", err)
	}
}

func init() {
	util.SetViperCfg(string(whopperutil.WhopperEngineParser), func() {
		// set config defaults
		viper.SetDefault("Port", 50052)
		viper.SetDefault("DaprStoreName", "statestore")
		viper.SetDefault("LogLevel", util.Debug)
		// set flags
		pflag.Bool("test", false, "testmode")
	})
}
