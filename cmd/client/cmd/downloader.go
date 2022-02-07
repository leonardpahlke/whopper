// Copyright 2022 Leonard Vincent Simon Pahlke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"time"
	"whopper/pkg/api"
	"whopper/pkg/whopperutil"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.grpcHost, cfg.grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
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
