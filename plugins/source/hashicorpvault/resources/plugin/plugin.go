package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/hashicorpvault/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var (
	Version = "Development"
)

func HashicorpVault() *source.Plugin {
	return source.NewPlugin(
		"hacshicorpvault",
		Version,
		tables(),
		client.Configure,
	)
}
