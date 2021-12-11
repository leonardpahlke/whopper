package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/util"
	"climatewhopper/pkg/whopperutil"
	"context"
	"fmt"
	"net"

	language "cloud.google.com/go/language/apiv1"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// implementedAnalyzerServer is used to implement Analyze function.
type implementedAnalyzerServer struct {
	logger     *zap.SugaredLogger
	ctx        context.Context
	daprClient dapr.Client
	nlpClient  *language.Client
	api.UnimplementedAnalyzerServer
}

// Translate function extends gRPC translator server function and is used to translate a newspaper article text to english
func (s implementedAnalyzerServer) Analyze(ctx context.Context, in *api.AnalyzerRequest) (*api.AnalyzerResponse, error) {
	// TODO: check which analytics are enabled
	s.logger.Infow("start analyze", "text length", len(in.Text))
	// get text sentiment by sentence
	sentiment, err := s.sentimentAnalysis(in.Text)
	if err != nil {
		return &api.AnalyzerResponse{
			Id:   in.Id,
			Head: &api.Head{Status: api.Status_ERROR, StatusMessage: errors.Wrap(err, "could not analyze sentiment").Error(), Timestamp: timestamppb.Now()},
		}, err
	}
	// get text entities
	entities, err := s.entitiesAnalysis(in.Text)
	if err != nil {
		return &api.AnalyzerResponse{
			Id:   in.Id,
			Head: &api.Head{Status: api.Status_ERROR, StatusMessage: errors.Wrap(err, "could not analyze entities").Error(), Timestamp: timestamppb.Now()},
		}, err
	}

	return &api.AnalyzerResponse{
		Id: in.Id,
		ArticleAnalytics: &api.ArticleAnalytics{
			AnalyzeEntitiesResponse:  entities,
			AnalyzeSentimentResponse: sentiment,
		},
		Head: &api.Head{Status: api.Status_OK, StatusMessage: "completed text analysis", Timestamp: timestamppb.Now()},
	}, nil
}

func main() {
	logger := util.GetLogger(util.MatchLogLevel(util.WrapLogLevel(viper.GetString("LogLevel"))))
	ctx := context.Background()

	// create dapr daprClient
	daprClient, err := dapr.NewClient()
	if err != nil {
		logger.Panicw("could not create dapr client", "error", err)
	}
	defer daprClient.Close()

	// Creates a gcp ML Client
	gcpNlpClient, err := language.NewClient(ctx)
	if err != nil {
		logger.Fatalw("could not create a client", "error", err)
	}
	defer gcpNlpClient.Close()

	// net listen
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("Port")))
	if err != nil {
		logger.Fatalw("failed to listen", "error", err)
	}

	// create new gRPC server
	s := grpc.NewServer()

	// Register server
	api.RegisterAnalyzerServer(s, &implementedAnalyzerServer{
		daprClient: daprClient,
		nlpClient:  gcpNlpClient,
		ctx:        ctx,
		logger:     util.GetLogger(zap.DebugLevel),
	})
	logger.Infow("server is listening", "address", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalw("failed to server", "error", err)
	}
}

// This method is used to run sentiment analysis
func (s *implementedAnalyzerServer) sentimentAnalysis(text string) (*languagepb.AnalyzeSentimentResponse, error) {
	s.logger.Debugw("run sentiment analysis", "text length", text)
	sentiment, err := s.nlpClient.AnalyzeSentiment(s.ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to analyze sentiment of text")
	}

	s.logger.Debugw("run sentiment analysis completed",
		"sentences", len(sentiment.Sentences),
		"language", sentiment.Language,
		"document sentiment score", sentiment.DocumentSentiment.Score,
		"sentiment is positive", sentiment.DocumentSentiment.Score >= 0,
	)

	return sentiment, nil
}

// This method is used to run entities analysis
func (s *implementedAnalyzerServer) entitiesAnalysis(text string) (*languagepb.AnalyzeEntitiesResponse, error) {
	s.logger.Debugw("run sentiment analysis", "text length", text)
	entities, err := s.nlpClient.AnalyzeEntities(s.ctx, &languagepb.AnalyzeEntitiesRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to analyze entities of text")
	}

	s.logger.Debugw("run entities analysis completed",
		"number of entities", len(entities.Entities),
		"language", entities.Language,
	)

	return entities, nil
}

func init() {
	util.SetViperCfg(string(whopperutil.WhopperEngineAnalyzer), func() {
		// set config defaults
		viper.SetDefault("Port", 50054)
		viper.SetDefault("DaprStoreName", "statestore")
		viper.SetDefault("LogLevel", util.Debug)
		// TODO: add configuration which can enable certain analysis
		// set flags
		pflag.Bool("test", false, "testmode")
	})
}
