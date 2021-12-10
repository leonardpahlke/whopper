package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/newsparser"
	"climatewhopper/pkg/util"
	"climatewhopper/pkg/whopperutil"
	"context"
	"fmt"
	"net"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// implementedDiscoveryServer is used to implement Discovery function.
type implementedDiscoveryServer struct {
	logger *zap.SugaredLogger
	api.UnimplementedDiscovererServer
}

func main() {
	// TODO: think about refactore this part of the code as it will reappear almost the same in all of the servers
	logger := util.GetLogger(util.MatchLogLevel(util.WrapLogLevel(viper.GetString("LogLevel"))))

	// net listen
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("Port")))
	if err != nil {
		logger.Fatalw("failed to listen", "error", err)
	}

	// create new gRPC server
	s := grpc.NewServer()
	// Register server
	api.RegisterDiscovererServer(s, &implementedDiscoveryServer{
		logger: util.GetLogger(util.MatchLogLevel(util.WrapLogLevel(viper.GetString("LogLevel")))),
	})
	logger.Infow("server is listening", "address", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalw("failed to server", "error", err)
	}
}

func init() {
	util.SetViperCfg(string(whopperutil.WhopperControllerDiscoverer), func() {
		// set config defaults
		viper.SetDefault("Port", 50055)
		viper.SetDefault("LogLevel", util.Debug)
		// set flags
		pflag.Bool("test", false, "testmode")
	})
}

func (s *implementedDiscoveryServer) Discover(ctx context.Context, in *api.DiscovererRequest) (*api.DiscovererResponse, error) {
	unprocessedArticles, err := newsparser.BatchDiscovery(in.Info)
	if err != nil {
		s.logger.Errorw("could not run batch discovery")
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
