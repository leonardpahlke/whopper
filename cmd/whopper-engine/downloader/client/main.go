package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/util"
	"context"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:39429"
)

func main() {
	// create logger
	logger := util.GetLogger(util.MatchLogLevel("DEBUG"))
	logger.Infow("Start downloader client", "address", address)

	// set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Fatalw("did not connect", "error", err)
	}
	defer conn.Close()
	c := api.NewDownloaderClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Download(ctx, &api.DownloadRequest{
		Id:  "sample-id",
		Url: "https://example.com",
	})
	if err != nil {
		logger.Fatalw("could not perform download request", "error", err)
	}
	logger.Infow("received response", "response", r, "response data length", len(r.Data))
}
