package codegen

import (
	"github.com/digitalocean/godo"
)

var emptyString = ""

var resources = []*Resource{
	{
		Service:             "account",
		Struct:              &godo.Account{},
		NewFunction:         godo.NewClient,
		SkipFetch:           true,
		PreResourceResolver: "accountGet",
		Multiplex:           &emptyString,
		ChildTable:          true,
	},
}

func Resources() []*Resource {

	for _, resource := range resources {
		resource.Template = "resource_get"
		resource.MockTemplate = "resource_get_mock"
	}

	return resources
}
