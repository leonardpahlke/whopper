package cmd

import (
	"climatewhopper/pkg/whopperutil"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var hubCmd = &cobra.Command{
	Use:          string(whopperutil.WhopperControllerHub),
	Short:        fmt.Sprintf("Send a request to the whopper-controller %s server", whopperutil.WhopperControllerHub),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runHubClient(cfg)
	},
}

// hub client function
func runHubClient(cfg *clientConfig) error {
	cfg.logger.Infof("start client %s", whopperutil.WhopperControllerHub)
	// TODO: client handler
	return errors.New("not implemented yet")
}

func init() {
	rootCmd.AddCommand(hubCmd)
}
