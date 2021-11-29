package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
}

// TODO: whopper-engine downloader

// handle gRPC request to download website resources

func init() {
	// load configuration
	viper.SetConfigName("we-downloader")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/climate-whopper/configs")
	viper.AddConfigPath("$HOME/climate-whopper/configs")
	viper.AddConfigPath("$HOME/Project/github/climate-whopper/configs")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// NOTE: defaults?
}
