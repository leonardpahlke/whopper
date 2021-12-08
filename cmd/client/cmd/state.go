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

var stateCfg = stateConfig{
	stateStoreName: "statestore",
	key:            "",
}
var stateCmd = &cobra.Command{
	Use:              "state",
	Short:            "Used to interact with the state storage",
	SilenceUsage:     true,
	TraverseChildren: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cfg.logger.Infow("run state command", "storename", stateCfg.stateStoreName, "key", stateCfg.key)
	},
}

func init() {
	// register state command to root command
	rootCmd.AddCommand(stateCmd)
	// set flags to state command
	stateCmd.PersistentFlags().StringVar(&stateCfg.stateStoreName, "statestore", stateCfg.stateStoreName, "Name of the state storage")
	stateCmd.PersistentFlags().StringVarP(&stateCfg.key, "key", "k", "", "State storage key")
	// set flags to required
	stateCmd.MarkPersistentFlagRequired("key")
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
