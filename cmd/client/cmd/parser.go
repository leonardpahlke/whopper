package cmd

import (
	"climatewhopper/pkg/whopperutil"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(parserCmd)
}

var parserCmd = &cobra.Command{
	Use:          string(whopperutil.WhopperEngineParser),
	Short:        fmt.Sprintf("Send a request to the whopper-engine %s server", whopperutil.WhopperEngineParser),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runParserClient(cfg)
	},
}

// parser client function
func runParserClient(cfg *clientConfig) error {
	cfg.logger.Infof("start client %s", whopperutil.WhopperEngineParser)
	// TODO: client handler
	return errors.New("not implemented yet")
}
