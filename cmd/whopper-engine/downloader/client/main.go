package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/util"
	"context"
	"time"

	v1 "github.com/dapr/dapr/pkg/proto/common/v1"
	pb "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	address = "localhost:38889"
)

func main() {
	// create logger
	logger := util.GetLogger(util.MatchLogLevel("DEBUG"))
	logger.Infow("Start downloader client", "address", address)

	// set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Fatalw("did not connect", "error", err)
	}
	defer conn.Close()
	c := pb.NewAppCallbackClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	any := anypb.Any{}
	any.MarshalFrom(&api.DownloadRequest{
		Id:  "sample-id",
		Url: "https://example.com",
	})
	r, err := c.OnInvoke(ctx, &v1.InvokeRequest{
		Method: "Download",
		Data:   &any,
	})

	if err != nil {
		logger.Fatalw("could not perform download request", "error", err)
	}
	logger.Infow("received response", "response", r, "response data length", len(r.Data.Value))
}
