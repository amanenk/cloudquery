// Code generated by codegen; DO NOT EDIT.

package projects

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/digitalocean/godo"
	"github.com/pkg/errors"

	"github.com/cloudquery/plugin-sdk/schema"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_projects",
		Resolver: fetchProjects,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "owner_uuid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerUUID"),
			},
			{
				Name:     "owner_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("OwnerID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "purpose",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Purpose"),
			},
			{
				Name:     "environment",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Environment"),
			},
			{
				Name:     "is_default",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsDefault"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},

		Relations: []*schema.Table{
			Resources(),
		},
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {

	svc := meta.(*client.Client)

	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}

	done := false
	listFunc := func() error {
		data, resp, err := svc.Services.Projects.List(ctx, opt)
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

	for !done {
		err := client.ThrottleWrapper(ctx, svc, listFunc)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
