package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/hashicorpvault/resources/services/pki"
	"github.com/cloudquery/plugin-sdk/schema"
)

func tables() []*schema.Table {
	return []*schema.Table{
		pki.Roles(),
	}
}
