package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

var sqlResources = []*Resource{
	{
		SubService: "instances",
		Struct:     &sqladmin.DatabaseInstance{},
	},
}

func SqlResources() []*Resource {
	var resources []*Resource
	resources = append(resources, sqlResources...)

	for _, resource := range resources {
		resource.Service = "sql"
		resource.Template = "resource_list"
		resource.ListFunction = fmt.Sprintf(`c.Services.Sql.%s.List(c.ProjectId).PageToken(nextPageToken).Do()`, strcase.ToCamel(resource.SubService))
	}

	return resources
}
