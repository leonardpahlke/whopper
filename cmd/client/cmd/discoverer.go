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

var discovererRequest = &api.DiscovererRequest{
	Info: []*api.InDiscovererInfo{{
		Newspaper: "taz",
		Url:       "https://taz.de/!t5204208/",
		LatestId:  5816465,
	}},
}

func init() {
	rootCmd.AddCommand(discovererCmd)
	// TODO: add flags or something to update client request
	// discovererCmd.Flags().String("id", downloaderRequest.Id, "request identifier that is used")
	// discovererCmd.Flags().String("url", downloaderRequest.Url, "article url which will be downloaded")
}

var discovererCmd = &cobra.Command{
	Use:          string(whopperutil.WhopperControllerDiscoverer),
	Short:        fmt.Sprintf("Send a request to the whopper-controller %s server", whopperutil.WhopperControllerDiscoverer),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDiscovererClient(cfg)
	},
}

// discoverer client function
func runDiscovererClient(cfg *clientConfig) error {
	cfg.logger.Infof("start client %s", whopperutil.WhopperControllerDiscoverer)

	// TODO: think about refactoring this part of the code
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.grpcHost, cfg.grpcPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "could not connect using gRPC")
	}
	defer conn.Close()

	// Create a new downloader client
	c := api.NewDiscovererClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", string(whopperutil.WhopperControllerDiscoverer))

	// Send download request to server
	r, err := c.Discover(ctx, discovererRequest)
	if err != nil {
		return errors.Wrap(err, "could not perform a discover request")
	}

	// TODO: check head

	cfg.logger.Infow("received response", "message", r.Head.StatusMessage, "number of articles", len(r.Articles), "unprocessed articles", r.Articles)
	return nil
}
