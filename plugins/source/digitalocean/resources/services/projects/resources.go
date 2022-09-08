// Code generated by codegen; DO NOT EDIT.

package projects

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/digitalocean/godo"
	"github.com/pkg/errors"

	"github.com/cloudquery/plugin-sdk/schema"
)

func Resources() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_projects_resources",
		Resolver: fetchResources,
		Columns: []schema.Column{
			{
				Name:     "urn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URN"),
			},
			{
				Name:     "assigned_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssignedAt"),
			},
			{
				Name:     "links",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Links"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
		},
	}
}

func fetchResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(godo.Project)
	svc := meta.(*client.Client)

	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}

	done := false
	for !done {
		listFunc := func() error {
			data, resp, err := svc.Services.Projects.ListResources(ctx, p.ID, opt)
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
