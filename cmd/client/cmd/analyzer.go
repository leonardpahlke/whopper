package cmd

import (
	"climatewhopper/pkg/whopperutil"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(analyzerCmd)
	analyzerCmd.MarkFlagRequired("port")
}

var analyzerCmd = &cobra.Command{
	Use:          string(whopperutil.WhopperEngineAnalyzer),
	Short:        fmt.Sprintf("Send a request to the whopper-engine %s server", whopperutil.WhopperEngineAnalyzer),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAnalyzerClient(cfg)
	},
}

// analyzer client function
func runAnalyzerClient(cfg *clientConfig) error {
	cfg.logger.Infof("start client %s", whopperutil.WhopperEngineAnalyzer)
	// TODO: client handler
	return errors.New("not implemented yet")
}
