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

	v1 "github.com/dapr/dapr/pkg/proto/common/v1"
	dapr "github.com/dapr/go-sdk/client"
	pb "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// server is used to implement dapr callback server.
type server struct {
	logger     *zap.SugaredLogger
	daprClient dapr.Client
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
	pb.RegisterAppCallbackServer(s, &server{
		daprClient: client,
		logger:     util.GetLogger(zap.DebugLevel),
	})

	logger.Infow("server is listening", "address", lis.Addr())

	if err := s.Serve(lis); err != nil {
		logger.Fatalw("failed to serve", "error", err)
	}
}

// This method gets invoked when a remote service has called the app through Dapr
// The payload carries a Method to identify the method, a set of metadata properties and an optional payload
// see: https://docs.dapr.io/developing-applications/integrations/grpc-integration/
func (s *server) OnInvoke(ctx context.Context, in *v1.InvokeRequest) (*v1.InvokeResponse, error) {
	s.logger.Infow("Server Invoke", "in", *in)
	var response *api.DownloadResponse
	downloadRequest := api.DownloadRequest{}
	err := in.Data.UnmarshalTo(downloadRequest.ProtoReflect().Interface())
	if err != nil {
		s.logger.Fatalw("could not unmarshal request", "input data", string(in.Data.Value))
	}

	if in.Method != "Download" {
		s.logger.Fatalw("unrecognized invoke method", "input method", string(in.Method))
	}

	response, err = s.Download(ctx, &downloadRequest)

	return &v1.InvokeResponse{
		ContentType: "text/plain; charset=UTF-8",
		Data:        &any.Any{Value: response.GetData()},
	}, err
}

// // Dapr will call this method to get the list of topics the app wants to subscribe to. In this example, we are telling Dapr
// // To subscribe to a topic named TopicA
// func (s *server) ListTopicSubscriptions(ctx context.Context, in *empty.Empty) (*pb.ListTopicSubscriptionsResponse, error) {
// 	return &pb.ListTopicSubscriptionsResponse{
// 		Subscriptions: []*pb.TopicSubscription{
// 			{Topic: "TopicA"},
// 		},
// 	}, nil
// }

// // Dapr will call this method to get the list of bindings the app will get invoked by. In this example, we are telling Dapr
// // To invoke our app with a binding named storage
// func (s *server) ListInputBindings(ctx context.Context, in *empty.Empty) (*pb.ListInputBindingsResponse, error) {
// 	return &pb.ListInputBindingsResponse{
// 		Bindings: []string{"storage"},
// 	}, nil
// }

// // This method gets invoked every time a new event is fired from a registered binding. The message carries the binding name, a payload and optional metadata
// func (s *server) OnBindingEvent(ctx context.Context, in *pb.BindingEventRequest) (*pb.BindingEventResponse, error) {
// 	fmt.Println("Invoked from binding")
// 	return &pb.BindingEventResponse{}, nil
// }

// // This method is fired whenever a message has been published to a topic that has been subscribed. Dapr sends published messages in a CloudEvents 0.3 envelope.
// func (s *server) OnTopicEvent(ctx context.Context, in *pb.TopicEventRequest) (*pb.TopicEventResponse, error) {
// 	fmt.Println("Topic message arrived")
// 	return &pb.TopicEventResponse{}, nil
// }

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
