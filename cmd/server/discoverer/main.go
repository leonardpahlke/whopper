package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/newsparser"
	"climatewhopper/pkg/util"
	"climatewhopper/pkg/whopperutil"
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/foolin/pagser"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// implementedDiscoveryServer is used to implement Discovery function.
type implementedDiscoveryServer struct {
	logger *zap.SugaredLogger
	api.UnimplementedDiscovererServer
}

func main() {
	// TODO: think about refactore this part of the code as it will reappear almost the same in all of the servers
	logger := util.GetLogger(util.MatchLogLevel(util.WrapLogLevel(viper.GetString("LogLevel"))))

	// net listen
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("Port")))
	if err != nil {
		logger.Fatalw("failed to listen", "error", err)
	}

	// create new gRPC server
	s := grpc.NewServer()
	// Register server
	api.RegisterDiscovererServer(s, &implementedDiscoveryServer{
		logger: util.GetLogger(util.MatchLogLevel(util.WrapLogLevel(viper.GetString("LogLevel")))),
	})
	logger.Infow("server is listening", "address", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalw("failed to server", "error", err)
	}
}

func init() {
	util.SetViperCfg(string(whopperutil.WhopperControllerDiscoverer), func() {
		// set config defaults
		viper.SetDefault("Port", 50055)
		viper.SetDefault("LogLevel", util.Debug)
		// set flags
		pflag.Bool("test", false, "testmode")
	})
}

func (s *implementedDiscoveryServer) Discover(ctx context.Context, in *api.DiscovererRequest) (*api.DiscovererResponse, error) {
	// This client is used to parse & discover unprocessed articles from newspaper websites
	p := pagser.New()

	unprocessedArticles := []*api.ArticleHead{}
	// loop over all newspapers that have been requested to be discovered
	// TODO: run discovery calls in concurrent (go routine)
	for _, e := range in.Info {
		switch e.Newspaper {
		// TODO: define newaspaper references with a type
		case "taz":
			{
				articles, err := newsparser.TazDiscoverer(p, e.Url)
				if err != nil {
					s.logger.Errorw("error running whopperutil.TazDiscoverer", "request  info", e)
					return &api.DiscovererResponse{
						Articles: unprocessedArticles,
						Head: &api.Head{
							Status:        api.Status_ERROR,
							StatusMessage: "received error during article discovery",
							Timestamp:     timestamppb.Now(),
						},
					}, err
				}
				s.logger.Infow("received taz discovery response", "amount of articles", len(articles.Articles))
				for _, a := range articles.Articles {
					// If the ID is not latestID the article has not been processed yet
					if a.ID != e.LatestId {
						unprocessedArticles = append(unprocessedArticles, &api.ArticleHead{
							Id:          a.ID,
							Url:         a.URL,
							ReleaseDate: a.Date,
							Title:       a.Title,
							Subtitle:    a.SubTitle,
							Description: a.Description,
							Category:    articles.Category,
						})
					} else {
						// The website structures articles in order, this allows to break if the ID matches -- all articles after that have already been processed
						break
					}
				}
			}
		default:
			err := errors.New("unrecognized input identifier")
			s.logger.Errorw("could not match newspaper reference to parser", "input information", in, "error", err)
			return nil, err
		}
	}

	return &api.DiscovererResponse{
		Articles: unprocessedArticles,
		Head: &api.Head{
			Status:        api.Status_OK,
			StatusMessage: fmt.Sprintf("successfully discovered new articles, count: %d", len(unprocessedArticles)),
			Timestamp:     timestamppb.Now(),
		},
	}, nil
}
