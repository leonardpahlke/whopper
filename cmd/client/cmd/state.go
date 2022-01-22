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

	dapr "github.com/dapr/go-sdk/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// This command is used to simplify how to interfact with the dapr statestorage
// The state command has subcommands to interact with the statestore

type stateConfig struct {
	stateStoreName string
	key            string
}

var (
	stateCfg = stateConfig{
		stateStoreName: "statestore",
		key:            "",
	}
	stateCmd = &cobra.Command{
		Use:              "state",
		Short:            "Used to interact with the state storage",
		SilenceUsage:     true,
		TraverseChildren: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cfg.logger.Infow("run state command", "storename", stateCfg.stateStoreName, "key", stateCfg.key)
		},
	}
)

func init() {
	// register state command to root command
	rootCmd.AddCommand(stateCmd)
	// set flags to state command
	stateCmd.PersistentFlags().StringVar(&stateCfg.stateStoreName, "statestore", stateCfg.stateStoreName, "Name of the state storage")
	stateCmd.PersistentFlags().StringVarP(&stateCfg.key, "key", "k", "", "State storage key")
	// set flags to required
	if err := stateCmd.MarkPersistentFlagRequired("key"); err != nil {
		panic(err)
	}
	// register state sub commands
	stateCmd.AddCommand(getStateCmd)
}

var getStateCmd = &cobra.Command{
	Use:          "get",
	Short:        "Retrieve data from state storage",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGetStateClient(cfg)
	},
}

// get state command
func runGetStateClient(cfg *clientConfig) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.grpcHost, cfg.grpcPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "could not connect using gRPC")
	}
	defer conn.Close()

	// create dapr client
	client := dapr.NewClientWithConnection(conn)
	if err != nil {
		cfg.logger.Panicw("could not create dapr client", "error", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// request state from dapr state storage
	state, err := client.GetState(ctx, stateCfg.stateStoreName, stateCfg.key)
	if err != nil {
		return errors.Wrap(err, "could not request state resource")
	}

	cfg.logger.Infow("received response from state storage", "state", *state)
	return nil
}
