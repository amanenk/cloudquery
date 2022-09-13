// Code generated by codegen; DO NOT EDIT.

package firewalls

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/digitalocean/godo"
)

func Firewalls() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_firewalls",
		Resolver: fetchFirewalls,
		Columns: []schema.Column{
			{
				Name:     "droplet_ids",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("DropletIDs"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "inbound_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InboundRules"),
			},
			{
				Name:     "outbound_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OutboundRules"),
			},
			{
				Name:     "droplet_i_ds",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("DropletIDs"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "created",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "pending_changes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingChanges"),
			},
		},
	}
}

func fetchFirewalls(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {

	svc := meta.(*client.Client)

	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}

	done := false
	listFunc := func() error {
		data, resp, err := svc.Services.Firewalls.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's data to our result channel
		res <- data
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			done = true
			return nil
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.Page = page + 1
		return nil
	}

	for !done {
		err := client.ThrottleWrapper(ctx, svc, listFunc)
		if err != nil {
			return err
		}
	}
	return nil
}
