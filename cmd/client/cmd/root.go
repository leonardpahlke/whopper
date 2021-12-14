package cmd

import (
	"climatewhopper/pkg/util"
	"climatewhopper/pkg/whopperutil"
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var cfg = &clientConfig{
	grpcPort: 0,
	grpcHost: "localhost",
	logger:   util.GetLogger(zap.InfoLevel),
}

type clientConfig struct {
	grpcPort int
	grpcHost string
	logger   *zap.SugaredLogger
}

var rootCmd = &cobra.Command{
	Use:     "wclient",
	Short:   "Whopper generic gRPC client CLI",
	Long:    fmt.Sprintf("Whopper controller and whopper engine gRPC client CLI which is used to send a gRPC request to one of the servers %v", whopperutil.WhopperServers),
	Example: "wclient [command] --port 54534 --host localhost",
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		cfg.logger.Infow("client request completed", "args", args)
	},
	TraverseChildren: false,
	SilenceUsage:     true,
}

// Execute executes the ci-reporter root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// set and parse flags
	rootCmd.PersistentFlags().StringVar(&cfg.grpcHost, "dapr-host", "localhost", "Specify where to reach dapr by providing an ipv4")
	rootCmd.PersistentFlags().IntVarP(&cfg.grpcPort, "dapr-grpc-port", "p", 0, "Dapr gRPC port")
	// mark required flags
	if err := rootCmd.MarkPersistentFlagRequired("port"); err != nil {
		fmt.Println("Info: required flag 'port' is not set")
	}
}
