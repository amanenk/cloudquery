package account

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
)

func TestIntegrationAccount(t *testing.T) {
	client.DOTestHelper(t, Account())
}
