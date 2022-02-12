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

func main() {}

// // implementedAnalyzerServer is used to implement Analyze function.
// type implementedAnalyzerServer struct {
// 	clients *whopperutil.WhopperServerClients
// 	config  *AnalyzerServerConfig
// 	analyzer.UnimplementedAnalyzerServer
// }

// // AnalyzerServerConfig configuration which will be injected as dapr sidecar information
// type AnalyzerServerConfig struct {
// 	Spec struct {
// 		Template struct {
// 			Metadata struct {
// 				Annotations struct {
// 					AppID       string `mapstructure:"dapr.io/app-id"`
// 					AppProtocol string `mapstructure:"dapr.io/app-protocol"`
// 					AppPort     string `mapstructure:"dapr.io/app-port"`
// 					LogLevel    string `mapstructure:"dapr.io/app-level"`
// 				} `mapstructure:"annotations"`
// 			} `mapstructure:"metadata"`
// 		} `mapstructure:"template"`
// 	} `mapstructure:"spec"`
// }

// func main() {
// 	// parse configuration file with viper
// 	config := AnalyzerServerConfig{}
// 	whopperutil.SetViperCfg(fmt.Sprintf("dapr-%s-config", whopperutil.WhopperEngineAnalyzer), func() {}, &config)

// 	// create clients for server
// 	clients := whopperutil.GetWhopperClient(whopperutil.WhopperServerClientIn{
// 		LogLevel:        config.Spec.Template.Metadata.Annotations.LogLevel,
// 		SetDaprClient:   true,
// 		SetGcpNlpClient: true,
// 	})
// 	// close clients after server shutsdown
// 	defer clients.Close()

// 	// create grpc server (without starting it yet)
// 	serverInit, err := whopperutil.CreateGrpcServer(config.Spec.Template.Metadata.Annotations.AppPort)
// 	if err != nil {
// 		clients.Logger.Fatalw("could not create plain grpc server", "error", err)
// 	}

// 	// register grpc translator server
// 	analyzer.RegisterAnalyzerServer(serverInit.Server, &implementedAnalyzerServer{clients: &clients, config: &config})

// 	// start listening on port
// 	serverInit.StartListening(clients.Logger)
// }

// // Translate function extends gRPC translator server function and is used to translate a newspaper article text to english
// func (s implementedAnalyzerServer) Analyze(ctx context.Context, in *analyzer.AnalyzerRequest) (*analyzer.AnalyzerResponse, error) {
// 	s.clients.Logger.Infow("start analyze", "text length", len(in.Text))
// 	// get text sentiment by sentence
// 	sentiment, err := s.sentimentAnalysis(in.Text)
// 	if err != nil {
// 		return &analyzer.AnalyzerResponse{
// 			Id:   in.Id,
// 			Head: &analyzer.Head{Status: analyzerStatus_ERROR, StatusMessage: errors.Wrap(err, "could not analyze sentiment").Error(), Timestamp: timestamppb.Now()},
// 		}, err
// 	}
// 	// get text entities
// 	entities, err := s.entitiesAnalysis(in.Text)
// 	if err != nil {
// 		return &analyzer.AnalyzerResponse{
// 			Id:   in.Id,
// 			Head: &analyzer.Head{Status: analyzer.Status_ERROR, StatusMessage: errors.Wrap(err, "could not analyze entities").Error(), Timestamp: timestamppb.Now()},
// 		}, err
// 	}

// 	return &analyzer.AnalyzerResponse{
// 		Id: in.Id,
// 		ArticleAnalytics: &analyzer.ArticleAnalytics{
// 			AnalyzeEntitiesResponse:  entities,
// 			AnalyzeSentimentResponse: sentiment,
// 		},
// 		Head: &analyzer.Head{Status: analyzerStatus_OK, StatusMessage: "completed text analysis", Timestamp: timestamppb.Now()},
// 	}, nil
// }

// //
// // Helper Methods
// //

// // This method is used to run sentiment analysis
// func (s *implementedAnalyzerServer) sentimentAnalysis(text string) (*languagepb.AnalyzeSentimentResponse, error) {
// 	s.clients.Logger.Debugw("run sentiment analysis", "text length", text)
// 	sentiment, err := s.clients.GcpLanguageClient.AnalyzeSentiment(s.clients.Ctx, &languagepb.AnalyzeSentimentRequest{
// 		Document: &languagepb.Document{
// 			Source: &languagepb.Document_Content{
// 				Content: text,
// 			},
// 			Type: languagepb.Document_PLAIN_TEXT,
// 		},
// 		EncodingType: languagepb.EncodingType_UTF8,
// 	})
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to analyze sentiment of text")
// 	}

// 	s.clients.Logger.Debugw("run sentiment analysis completed",
// 		"sentences", len(sentiment.Sentences),
// 		"language", sentiment.Language,
// 		"document sentiment score", sentiment.DocumentSentiment.Score,
// 		"sentiment is positive", sentiment.DocumentSentiment.Score >= 0,
// 	)

// 	return sentiment, nil
// }

// // This method is used to run entities analysis
// func (s *implementedAnalyzerServer) entitiesAnalysis(text string) (*languagepb.AnalyzeEntitiesResponse, error) {
// 	s.clients.Logger.Debugw("run sentiment analysis", "text length", text)
// 	entities, err := s.clients.GcpLanguageClient.AnalyzeEntities(s.clients.Ctx, &languagepb.AnalyzeEntitiesRequest{
// 		Document: &languagepb.Document{
// 			Source: &languagepb.Document_Content{
// 				Content: text,
// 			},
// 			Type: languagepb.Document_PLAIN_TEXT,
// 		},
// 		EncodingType: languagepb.EncodingType_UTF8,
// 	})
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to analyze entities of text")
// 	}

// 	s.clients.Logger.Debugw("run entities analysis completed",
// 		"number of entities", len(entities.Entities),
// 		"language", entities.Language,
// 	)

// 	return entities, nil
// }
