package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/whopperutil"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// implementedDownloaderServer is used to implement Download function.
type implementedDownloaderServer struct {
	clients *whopperutil.WhopperServerClients
	config  *DownloaderServerConfig
	api.UnimplementedDownloaderServer
}

// DownloaderServerConfig configuration which will be injected as dapr sidecar information
type DownloaderServerConfig struct {
	Spec struct {
		Template struct {
			Metadata struct {
				DaprStoreName string `mapstructure:"daprStoreName"`
				Annotations   struct {
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
	config := DownloaderServerConfig{}
	whopperutil.SetViperCfg(fmt.Sprintf("dapr-%s-config", whopperutil.WhopperEngineDownloader), func() {}, &config)

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
	api.RegisterDownloaderServer(serverInit.Server, &implementedDownloaderServer{clients: &clients, config: &config})

	// start listening on port
	serverInit.StartListening(clients.Logger)
}

//
// Server Methods
//

// Download implements api.DownloaderServer
func (s *implementedDownloaderServer) Download(ctx context.Context, in *api.DownloadRequest) (*api.DownloadResponse, error) {
	s.clients.Logger.Infow("received download invoke", "id", in.Id)

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

	// read http body
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
	s.clients.Logger.Debugw("write body to state storage", "body length", len(body))

	// store body to database
	err = s.clients.DaprClient.SaveState(ctx, viper.GetString(s.config.Spec.Template.Metadata.DaprStoreName), in.Id, body)
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

	// create grpc response
	return &api.DownloadResponse{
		Id: in.Id,
		ArticleFooter: &api.ArticleFooter{
			RawArticleText: string(body),
		},
		Head: &api.Head{Status: api.Status_OK, StatusMessage: "data downloaded and stored", Timestamp: timestamppb.Now()},
	}, nil
}
