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
	"climatewhopper/pkg/whopperutil"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(analyzerCmd)
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
