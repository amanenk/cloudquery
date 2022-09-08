// Code generated by codegen; DO NOT EDIT.

package monitoring

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/digitalocean/godo"
	"github.com/pkg/errors"

	"github.com/cloudquery/plugin-sdk/schema"
)

func AlertPolicies() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_monitoring_alert_policies",
		Resolver: fetchAlertPolicies,
		Columns: []schema.Column{
			{
				Name:     "uuid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UUID"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "compare",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Compare"),
			},
			{
				Name:     "value",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("Value"),
			},
			{
				Name:     "window",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Window"),
			},
			{
				Name:     "entities",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Entities"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "alerts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Alerts"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
		},
	}
}

func fetchAlertPolicies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {

	svc := meta.(*client.Client)

	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}

	done := false
	for !done {
		listFunc := func() error {
			data, resp, err := svc.Services.Monitoring.ListAlertPolicies(ctx, opt)
			if err != nil {
				return errors.WithStack(err)
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
				return errors.WithStack(err)
			}
			// set the page we want for the next request
			opt.Page = page + 1
			return nil
		}

		err := client.ThrottleWrapper(ctx, svc, listFunc)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
