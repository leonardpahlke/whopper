package cmd

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/whopperutil"
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var downloaderRequest = &api.DownloadRequest{
	Id:  "example",
	Url: "https://example.com",
}

func init() {
	rootCmd.AddCommand(downloaderCmd)
	downloaderCmd.Flags().String("id", downloaderRequest.Id, "request identifier that is used")
	downloaderCmd.Flags().String("url", downloaderRequest.Url, "article url which will be downloaded")
}

var downloaderCmd = &cobra.Command{
	Use:          string(whopperutil.WhopperEngineDownloader),
	Short:        fmt.Sprintf("Send a request to the whopper-engine %s server", whopperutil.WhopperEngineDownloader),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDownloaderClient(cfg)
	},
}

// downloader client function
func runDownloaderClient(cfg *clientConfig) error {
	cfg.logger.Infof("start client %s", whopperutil.WhopperEngineDownloader)

	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.grpcHost, cfg.grpcPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "could not connect using gRPC")
	}
	defer conn.Close()

	// Create a new downloader client
	c := api.NewDownloaderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", string(whopperutil.WhopperEngineDownloader))

	// Send download request to server
	r, err := c.Download(ctx, downloaderRequest)
	if err != nil {
		return errors.Wrap(err, "could not perform a download request")
	}

	// TODO: check head

	cfg.logger.Infow("received response", "response id", r.Id, "response data length", len(r.ArticleFooter.RawArticleText))
	return nil
}
