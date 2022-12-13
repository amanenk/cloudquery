// Code generated by codegen2; DO NOT EDIT.
package marketplace

import (
	"encoding/json"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createPrivateStore(router *mux.Router) error {
	var item armmarketplace.PrivateStoreClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/providers/Microsoft.Marketplace/privateStores", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	return nil
}

func TestPrivateStore(t *testing.T) {
	client.MockTestHelper(t, PrivateStore(), createPrivateStore)
}