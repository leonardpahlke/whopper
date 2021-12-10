package whopperutil

// Package with whopper related utility

// WhopperServerName type used for gRPC servers of this project
type WhopperServerName string

var (
	// WhopperControllerDiscoverer discoverer server name
	WhopperControllerDiscoverer WhopperServerName = "discoverer"
	// WhopperControllerHub hub server name
	WhopperControllerHub WhopperServerName = "hub"
	// WhopperEngineDownloader downloader server name
	WhopperEngineDownloader WhopperServerName = "downloader"
	// WhopperEngineParser parser server name
	WhopperEngineParser WhopperServerName = "parser"
	// WhopperEngineTranslator translator server name
	WhopperEngineTranslator WhopperServerName = "translator"
	// WhopperEngineAnalyzer analyzer server name
	WhopperEngineAnalyzer WhopperServerName = "analyzer"

	// WhopperServers available whopper servers
	WhopperServers = []WhopperServerName{WhopperControllerDiscoverer, WhopperControllerHub, WhopperEngineDownloader, WhopperEngineParser, WhopperEngineTranslator, WhopperEngineAnalyzer}
)
