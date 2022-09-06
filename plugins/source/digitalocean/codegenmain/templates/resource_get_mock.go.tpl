package {{.Service}}

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
)

func TestIntegration{{.Service | ToCamel}}(t *testing.T) {
	client.DOTestHelper(t, {{.Service | ToCamel}}())
}
