package main

import (
	"climatewhopper/pkg/api"
	"climatewhopper/pkg/util"
	"context"

	dapr "github.com/dapr/go-sdk/client"
	"google.golang.org/grpc"
)

const (
	address = "localhost:40235"
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

	// create dapr client
	client := dapr.NewClientWithConnection(conn)
	// if err != nil {
	// 	logger.Panicw("could not create dapr client", "error", err)
	// }
	defer client.Close()

	// ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	// defer cancel()
	ctx := context.Background()

	// data, err := proto.Marshal(&api.DownloadRequest{
	// 	Id:  "sample-id",
	// 	Url: "https://example.com",
	// })
	// if err != nil {
	// 	logger.Panicw("data could not be transformed to bytes", "error", err)
	// }
	out, err := client.InvokeMethodWithCustomContent(ctx, "downloader", "Download", "test", "text/plain; charset=UTF-8", &api.DownloadRequest{
		Id:  "sample-id",
		Url: "https://example.com",
	})

	// c := pb.NewAppCallbackClient(conn)

	// // Contact the server and print out its response.

	// any := anypb.Any{}
	// any.MarshalFrom(&api.DownloadRequest{
	// 	Id:  "sample-id",
	// 	Url: "https://example.com",
	// })
	// r, err := c.OnInvoke(ctx, &v1.InvokeRequest{
	// 	Method: "Download",
	// 	Data:   &any,
	// })

	if err != nil {
		logger.Fatalw("could not perform download request", "error", err)
	}
	logger.Infow("received response", "response", string(out), "response data length", len(out))
}
