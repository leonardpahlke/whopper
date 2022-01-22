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

package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/whopperutil"
	"context"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/text/language"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Type used to implement the Translator grpc server
type implementedTranslatorServer struct {
	clients *whopperutil.WhopperServerClients
	config  *TranslatorServerConfig
	api.UnimplementedTranslatorServer
}

// TranslatorServerConfig configuration which will be injected as dapr sidecar information
type TranslatorServerConfig struct {
	Spec struct {
		Template struct {
			Metadata struct {
				DaprStoreName  string `mapstructure:"daprStoreName"`
				TargetLanguage string `mapstructure:"targetLanguage"`
				Annotations    struct {
					AppID       string `mapstructure:"dapr.io/app-id"`
					AppProtocol string `mapstructure:"dapr.io/app-protocol"`
					AppPort     string `mapstructure:"dapr.io/app-port"`
					LogLevel    string `mapstructure:"dapr.io/app-level"`
				} `mapstructure:"annotations"`
			} `mapstructure:"metadata"`
		} `mapstructure:"template"`
	} `mapstructure:"spec"`
}

func main() {
	// parse configuration file with viper
	config := TranslatorServerConfig{}
	whopperutil.SetViperCfg(fmt.Sprintf("dapr-%s-config", whopperutil.WhopperEngineTranslator), func() {}, &config)

	// create clients for server
	clients := whopperutil.GetWhopperClient(whopperutil.WhopperServerClientIn{
		LogLevel:              config.Spec.Template.Metadata.Annotations.LogLevel,
		SetGcpTranslateClient: true,
	})
	// close clients after server shutsdown
	defer clients.Close()

	// create grpc server (without starting it yet)
	serverInit, err := whopperutil.CreateGrpcServer(config.Spec.Template.Metadata.Annotations.AppPort)
	if err != nil {
		clients.Logger.Fatalw("could not create plain grpc server", "error", err)
	}

	// register grpc translator server
	api.RegisterTranslatorServer(serverInit.Server, &implementedTranslatorServer{clients: &clients, config: &config})

	// start listening on port
	serverInit.StartListening(clients.Logger)
}

//
// Server Methods
//

// Translate function extends gRPC translator server function and is used to translate a newspaper article text to english
func (s implementedTranslatorServer) Translate(ctx context.Context, in *api.TranslatorRequest) (*api.TranslatorResponse, error) {
	s.clients.Logger.Infow("Translate article text", "target language", s.config.Spec.Template.Metadata.TargetLanguage)

	// translate article text
	translatedText, err := s.translateText(ctx, s.config.Spec.Template.Metadata.TargetLanguage, in.Text)
	if err != nil {
		return &api.TranslatorResponse{
			Id:   in.Id,
			Head: &api.Head{Status: api.Status_ERROR, StatusMessage: errors.Wrap(err, "could not translate text").Error(), Timestamp: timestamppb.Now()},
		}, err
	}
	s.clients.Logger.Debugw("text has been translated successfully", "translated text length", len(translatedText), "original text length", len(in.Text))
	if translatedText == in.Text {
		return &api.TranslatorResponse{
			Id:             in.Id,
			TranslatedText: translatedText,
			Head:           &api.Head{Status: api.Status_WARNING, StatusMessage: "translated text is equal to the original text", Timestamp: timestamppb.Now()},
		}, nil
	}
	// TODO: check if article text should be stored in statestore, if not remove the dapr client
	return &api.TranslatorResponse{
		Id:             in.Id,
		TranslatedText: translatedText,
		Head:           &api.Head{Status: api.Status_OK, StatusMessage: "text has been translated", Timestamp: timestamppb.Now()},
	}, nil
}

//
// Helper methods
//

// translateText method uses GCP translate ML service to translate article text
func (s implementedTranslatorServer) translateText(ctx context.Context, targetLanguage, text string) (string, error) {
	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", errors.Wrap(err, "could not parse language reference")
	}

	resp, err := s.clients.GcpTranslateClient.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", errors.Wrap(err, "could not translate text")
	}
	if len(resp) == 0 {
		return "", errors.Wrap(err, "empty translate response")
	}
	return resp[0].Text, nil
}
