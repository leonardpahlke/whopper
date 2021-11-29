package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/util"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	commonv1pb "github.com/dapr/go-sdk/dapr/proto/common/v1"
	pb "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"
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
	pb.UnimplementedAppCallbackServer
}

// Download implements api.DownloaderServer
func (s *server) Download(ctx context.Context, in *api.DownloadRequest) (*api.DownloadResponse, error) {
	log.Printf("Received download invoke, id: %s", in.Id)

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

// This method gets invoked when a remote service has called the app through Dapr
// The payload carries a Method to identify the method, a set of metadata properties and an optional payload
// see: https://docs.dapr.io/developing-applications/integrations/grpc-integration/
func (s *server) OnInvoke(ctx context.Context, in *commonv1pb.InvokeRequest) (*commonv1pb.InvokeResponse, error) {
	var response *api.DownloadResponse
	downloadRequest := api.DownloadRequest{}
	err := in.Data.UnmarshalTo(downloadRequest.ProtoReflect().Interface())
	if err != nil {
		s.logger.Fatalw("could not unmarshal request", "input data", string(in.Data.Value))
	}

	if in.Method != "Download" {
		s.logger.Fatalw("unrecognized invoke method", "input method", string(in.Method))
	}

	response, err = s.UnimplementedDownloaderServer.Download(ctx, &downloadRequest)
	if err != nil {
		s.logger.Errorw("download error occurred", "error", err)
	}

	return &commonv1pb.InvokeResponse{
		ContentType: "text/plain; charset=UTF-8",
		Data:        &any.Any{Value: response.GetData()},
	}, nil
}

// Dapr will call this method to get the list of bindings the app will get invoked by. In this example, we are telling Dapr
// To invoke our app with a binding named storage
func (s *server) ListInputBindings(ctx context.Context, in *empty.Empty) (*pb.ListInputBindingsResponse, error) {
	return &pb.ListInputBindingsResponse{
		Bindings: []string{"storage"},
	}, nil
}

func init() {
	util.SetViperCfg("downloader", func() {
		// set config defaults
		viper.SetDefault("Port", 50051)
		viper.SetDefault("DaprStoreName", "statestore")
		viper.SetDefault("LogLevel", util.Debug)
		// set flags
		pflag.Bool("test", false, "testmode")
	})
}
