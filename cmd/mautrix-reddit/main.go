package main

import (
	"maunium.net/go/mautrix/bridgev2"
	"maunium.net/go/mautrix/bridgev2/matrix/mxmain"

	"github.com/yourusername/mautrix-reddit/pkg/connector"
)

// Information to find out exactly which commit the bridge was built from.
// These are filled at build time with the -X linker flag.
var (
	Tag       = "unknown"
	Commit    = "unknown"
	BuildTime = "unknown"
)

func main() {
	m := mxmain.BridgeMain{
		Name:        "mautrix-reddit",
		URL:         "https://github.com/yourusername/mautrix-reddit",
		Description: "A Matrix-Reddit bridge",
		Version:     "0.1.0",
	}
	
	m.PostInit = func() {
		m.Bridge.Network = &connector.RedditConnector{
			Config: &connector.Config{},
		}
	}
	
	m.Run()
}
