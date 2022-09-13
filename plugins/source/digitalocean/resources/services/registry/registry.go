// Code generated by codegen; DO NOT EDIT.

package registry

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Registry() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_registry",
		Resolver: fetchRegistry,
		Columns: []schema.Column{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "storage_usage_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("StorageUsageBytes"),
			},
			{
				Name:     "storage_usage_bytes_updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StorageUsageBytesUpdatedAt"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
		},

		Relations: []*schema.Table{
			Repositories(),
		},
	}
}

func fetchRegistry(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {

	svc := meta.(*client.Client)
	getFunc := func() error {
		response, _, err := svc.Services.Registry.Get(ctx)
		if err != nil {
			return err
		}
		res <- response
		return nil
	}

	err := client.ThrottleWrapper(ctx, svc, getFunc)
	if err != nil {
		return err
	}
	return nil
}
