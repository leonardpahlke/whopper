package cmd

import (
	"climatewhopper/pkg/whopperutil"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(translatorCmd)
}

var translatorCmd = &cobra.Command{
	Use:          string(whopperutil.WhopperEngineTranslator),
	Short:        fmt.Sprintf("Send a request to the whopper-engine %s server", whopperutil.WhopperEngineTranslator),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runTranslatorClient(cfg)
	},
}

// translator client function
func runTranslatorClient(cfg *clientConfig) error {
	cfg.logger.Infof("start client %s", whopperutil.WhopperEngineTranslator)
	// TODO: client handler
	return errors.New("not implemented yet")
}
