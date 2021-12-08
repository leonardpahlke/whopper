package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/util"
	"climatewhopper/pkg/whopperutil"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	logger     *zap.SugaredLogger
	daprClient dapr.Client
	api.UnimplementedDownloaderServer
}

// Download implements api.DownloaderServer
func (s *server) Download(ctx context.Context, in *api.DownloadRequest) (*api.DownloadResponse, error) {
	s.logger.Infow("received download invoke", "id", in.Id)

	// get website data
	resp, err := http.Get(in.Url)
	if err != nil {
		return &api.DownloadResponse{
			Id: in.Id,
			Head: &api.Head{
				Status:        api.Status_ERROR,
				StatusMessage: errors.Wrap(err, "http request error").Error(),
				Timestamp:     timestamppb.Now(),
			},
		}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &api.DownloadResponse{
			Id: in.Id,
			Head: &api.Head{
				Status:        api.Status_ERROR,
				StatusMessage: errors.Wrap(err, "could not read http response body").Error(),
				Timestamp:     timestamppb.Now(),
			},
		}, err
	}
	s.logger.Debugw("write body to state storage", "body length", len(body))
	// store data to database
	err = s.daprClient.SaveState(ctx, viper.GetString("DaprStoreName"), in.Id, body)
	if err != nil {
		return &api.DownloadResponse{
			Id: in.Id,
			Head: &api.Head{
				Status:        api.Status_ERROR,
				StatusMessage: errors.Wrap(err, "could not save data to remote storage").Error(),
				Timestamp:     timestamppb.Now(),
			},
		}, err
	}
	s.logger.Debugw("return OK response", "id", in.Id)
	state, err := s.daprClient.GetState(ctx, viper.GetString("DaprStoreName"), in.Id)
	if err != nil {
		return &api.DownloadResponse{
			Id: in.Id,
			Head: &api.Head{
				Status:        api.Status_ERROR,
				StatusMessage: errors.Wrap(err, "could not get saved data from remote storage").Error(),
				Timestamp:     timestamppb.Now(),
			},
		}, err
	}

	s.logger.Debugw("received state", "state", *state)
	// response
	return &api.DownloadResponse{
		Id:   in.Id,
		Data: body,
		Head: &api.Head{
			Status:        api.Status_OK,
			StatusMessage: "data downloaded and stored",
			Timestamp:     timestamppb.Now(),
		},
	}, nil
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
	api.RegisterDownloaderServer(s, &server{
		daprClient: client,
		logger:     util.GetLogger(zap.DebugLevel),
	})
	logger.Infow("server is listening", "address", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalw("failed to server", "error", err)
	}
}

func init() {
	util.SetViperCfg(string(whopperutil.WhopperEngineDownloader), func() {
		// set config defaults
		viper.SetDefault("Port", 50051)
		viper.SetDefault("DaprStoreName", "statestore")
		viper.SetDefault("LogLevel", util.Debug)
		// set flags
		pflag.Bool("test", false, "testmode")
	})
}
