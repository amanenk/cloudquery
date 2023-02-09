package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/hashicorpvault/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

// FIXME replace with real DSN
const sentryDSN = "https://xxx"

func main() {
	serve.Source(plugin.HashicorpVault(), serve.WithSourceSentryDSN(sentryDSN))
}
