package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/util"
	"climatewhopper/pkg/whopperutil"
	"context"
	"fmt"
	"net"

	"cloud.google.com/go/translate"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// implementedTranslatorServer is used to implement Translate function.
type implementedTranslatorServer struct {
	logger     *zap.SugaredLogger
	daprClient dapr.Client
	api.UnimplementedTranslatorServer
}

// Translate function extends gRPC translator server function and is used to translate a newspaper article text to english
func (s implementedTranslatorServer) Translate(ctx context.Context, in *api.TranslatorRequest) (*api.TranslatorResponse, error) {
	targetLanguage := viper.GetString("TargetLanguage")
	s.logger.Infow("Translate article text", "target language", targetLanguage)

	// translate article text
	translatedText, err := s.translateText(ctx, targetLanguage, in.Text)
	if err != nil {
		return &api.TranslatorResponse{
			Id:   in.Id,
			Head: &api.Head{Status: api.Status_ERROR, StatusMessage: errors.Wrap(err, "could not translate text").Error(), Timestamp: timestamppb.Now()},
		}, err
	}
	s.logger.Debugw("text has been translated successfully", "translated text length", len(translatedText), "original text length", len(in.Text))
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

func main() {
	logger := util.GetLogger(util.MatchLogLevel(util.WrapLogLevel(viper.GetString("LogLevel"))))

	// create dapr client
	client, err := dapr.NewClient()
	if err != nil {
		logger.Panicw("could not create dapr client", "error", err)
	}
	defer client.Close()

	// net listen
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("Port")))
	if err != nil {
		logger.Fatalw("failed to listen", "error", err)
	}

	// create new gRPC server
	s := grpc.NewServer()
	// Register server
	api.RegisterTranslatorServer(s, &implementedTranslatorServer{
		daprClient: client,
		logger:     util.GetLogger(zap.DebugLevel),
	})
	logger.Infow("server is listening", "address", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalw("failed to server", "error", err)
	}
}

// translateText method uses GCP translate ML service to translate article text
func (s implementedTranslatorServer) translateText(ctx context.Context, targetLanguage, text string) (string, error) {
	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", errors.Wrap(err, "could not parse language reference")
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", errors.Wrap(err, "could not translate text")
	}
	if len(resp) == 0 {
		return "", errors.Wrap(err, "empty translate response")
	}
	return resp[0].Text, nil
}

func init() {
	util.SetViperCfg(string(whopperutil.WhopperEngineTranslator), func() {
		// set config defaults
		viper.SetDefault("Port", 50053)
		viper.SetDefault("DaprStoreName", "statestore")
		viper.SetDefault("TargetLanguage", "en-US")
		viper.SetDefault("LogLevel", util.Debug)
		// set flags
		pflag.Bool("test", false, "testmode")
	})
}
