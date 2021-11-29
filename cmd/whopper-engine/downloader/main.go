package main

import (
	"climatewhopper/pkg/api"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	api.UnimplementedDownloaderServer
}

// Download implements api.DownloaderServer
func (s *server) Download(ctx context.Context, in *api.DownloadRequest) (*api.DownloadResponse, error) {
	log.Printf("Received download invoke, id: %s", in.Id)

	// get website data
	resp, err := http.Get(in.Url)
	if err != nil {
		return &api.DownloadResponse{
			Id:        in.Id,
			Newspaper: in.Newspaper,
			Head: &api.Head{
				Status:        api.Status_ERROR,
				StatusMessage: errors.Wrap(err, "http request error").Error(),
				Timestamp:     timestamppb.Now(),
			},
		}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &api.DownloadResponse{
			Id:        in.Id,
			Newspaper: in.Newspaper,
			Head: &api.Head{
				Status:        api.Status_ERROR,
				StatusMessage: errors.Wrap(err, "could not read http response body").Error(),
				Timestamp:     timestamppb.Now(),
			},
		}, err
	}

	// TODO: store data to remote storage
	// ...

	return &api.DownloadResponse{
		Id:        in.Id,
		Newspaper: in.Newspaper,
		Data:      body,
		DataPath:  "",
		Head: &api.Head{
			Status:        api.Status_OK,
			StatusMessage: "data downloaded and stored",
			Timestamp:     timestamppb.Now(),
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("Port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterDownloaderServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func init() {
	// set config meta
	viper.SetConfigName("downloader")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./climate-whopper/configs")
	viper.AddConfigPath("$HOME/climate-whopper/configs")
	viper.AddConfigPath("./configs")
	// set config defaults
	viper.SetDefault("Port", 50051)
	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
