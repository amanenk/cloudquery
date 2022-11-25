// Code generated by codegen; DO NOT EDIT.

package certificates

import (
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CertificatePackages() *schema.Table {
	return &schema.Table{
		Name:     "gandi_certificate_packages",
		Resolver: fetchCertificatePackages,
		Columns: []schema.Column{
			{
				Name:        "sharing_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveSharingID,
				Description: `The Sharing ID of the resource.`,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name_label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NameLabel"),
			},
			{
				Name:     "href",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Href"),
			},
			{
				Name:     "max_domains",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxDomains"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "type_label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TypeLabel"),
			},
			{
				Name:     "wildcard",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Wildcard"),
			},
		},
	}
}
