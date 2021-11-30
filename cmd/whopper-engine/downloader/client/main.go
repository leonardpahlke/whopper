package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/util"
	"context"
	"time"

	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:43667"
)

func main() {
	logger := util.GetLogger(zapcore.DebugLevel)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Fatalw("did not connect", "error", err)
	}
	defer conn.Close()
	c := api.NewDownloaderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "downloader")
	r, err := c.Download(ctx, &api.DownloadRequest{
		Id:  "example",
		Url: "https://example.com",
	})
	if err != nil {
		logger.Fatalw("could not invoke", "error", err)
	}

	logger.Infow("received response", "response id", r.Id, "response data length", len(r.Data))
}
