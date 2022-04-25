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

// var discovererRequest = &api.DiscovererRequest{
// 	Info: []*api.InDiscovererInfo{{
// 		Newspaper: "taz",
// 		Url:       "https://taz.de/!t5204208/",
// 		LatestId:  5816465,
// 	}},
// }

// func init() {
// 	rootCmd.AddCommand(discovererCmd)
// 	// add flags or something to update client request
// 	// discovererCmd.Flags().String("id", downloaderRequest.Id, "request identifier that is used")
// 	// discovererCmd.Flags().String("url", downloaderRequest.Url, "article url which will be downloaded")
// }

// var discovererCmd = &cobra.Command{
// 	Use:          string(whopperutil.WhopperControllerDiscoverer),
// 	Short:        fmt.Sprintf("Send a request to the whopper-controller %s server", whopperutil.WhopperControllerDiscoverer),
// 	SilenceUsage: true,
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		return runDiscovererClient(cfg)
// 	},
// }

// // discoverer client function
// func runDiscovererClient(cfg *clientConfig) error {
// 	cfg.logger.Infof("start client %s", whopperutil.WhopperControllerDiscoverer)

// 	// think about refactoring this part of the code
// 	// Set up a connection to the server.
// 	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.grpcHost, cfg.grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
// 	if err != nil {
// 		return errors.Wrap(err, "could not connect using gRPC")
// 	}
// 	defer conn.Close()

// 	// Create a new downloader client
// 	c := api.NewDiscovererClient(conn)

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
// 	defer cancel()
// 	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", string(whopperutil.WhopperControllerDiscoverer))

// 	// Send download request to server
// 	r, err := c.Discover(ctx, discovererRequest)
// 	if err != nil {
// 		return errors.Wrap(err, "could not perform a discover request")
// 	}

// 	// NOTE: check head

// 	cfg.logger.Infow("received response", "message", r.Head.StatusMessage, "number of articles", len(r.Articles), "unprocessed articles", r.Articles)
// 	return nil
// }
