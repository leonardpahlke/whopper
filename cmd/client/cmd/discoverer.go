package cmd

import (
	"climatewhopper/pkg/whopperutil"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(discovererCmd)
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
	// TODO: client handler
	return errors.New("not implemented yet")
}
