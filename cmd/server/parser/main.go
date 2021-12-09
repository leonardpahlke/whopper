package main

import (
	"climatewhopper/pkg/util"
	"climatewhopper/pkg/whopperutil"
	"encoding/json"
	"fmt"
	"log"
	"net"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/foolin/pagser"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

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
	// api.RegisterDownloaderServer(s, &server{
	// 	daprClient: client,
	// 	logger:     util.GetLogger(zap.DebugLevel),
	// })
	logger.Infow("server is listening", "address", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalw("failed to server", "error", err)
	}
}

func init() {
	util.SetViperCfg(string(whopperutil.WhopperEngineParser), func() {
		// set config defaults
		viper.SetDefault("Port", 50052)
		viper.SetDefault("DaprStoreName", "statestore")
		viper.SetDefault("LogLevel", util.Debug)
		// set flags
		pflag.Bool("test", false, "testmode")
	})
}

// testing a library

const rawPageHTML = `
<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>Pagser Title</title>
	<meta name="keywords" content="golang,pagser,goquery,html,page,parser,colly">
</head>

<body>
	<h1>H1 Pagser Example</h1>
	<div class="navlink">
		<div class="container">
			<ul class="clearfix">
				<li id=''><a href="/">Index</a></li>
				<li id='2'><a href="/list/web" title="web site">Web page</a></li>
				<li id='3'><a href="/list/pc" title="pc page">Pc Page</a></li>
				<li id='4'><a href="/list/mobile" title="mobile page">Mobile Page</a></li>
			</ul>
		</div>
	</div>
</body>
</html>
`

// PageData is a test struct
type PageData struct {
	Title    string   `pagser:"title"`
	Keywords []string `pagser:"meta[name='keywords']->attrSplit(content)"`
	H1       string   `pagser:"h1"`
	Navs     []struct {
		ID   int    `pagser:"->attrEmpty(id, -1)"`
		Name string `pagser:"a->text()"`
		URL  string `pagser:"a->attr(href)"`
	} `pagser:".navlink li"`
}

func test() {
	//New default config
	p := pagser.New()

	//data parser model
	var data PageData
	//parse html data
	err := p.Parse(&data, rawPageHTML)
	//check error
	if err != nil {
		log.Fatal(err)
	}

	//print data
	log.Printf("Page data json: \n-------------\n%v\n-------------\n", toJSON(data))
}

func toJSON(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "\t")
	return string(data)
}
