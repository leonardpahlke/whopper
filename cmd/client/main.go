package main

import (
	"climatewhopper/cmd/client/cmd"
	"climatewhopper/pkg/util"

	"go.uber.org/zap"
)

// Entry point for the wclient cli
// The cli uses cobra and therefore cobra project structuring conventions
// und the cmd/ folder the main root command (root.go) and sub commands are defined
func main() {
	if err := cmd.Execute(); err != nil {
		util.GetLogger(zap.InfoLevel).Fatalw("could not execute command", "error", err)
	}
}
