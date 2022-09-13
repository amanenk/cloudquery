// Code generated by codegen; DO NOT EDIT.

package images

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/digitalocean/godo"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_images",
		Resolver: fetchImages,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "distribution",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Distribution"),
			},
			{
				Name:     "slug",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Slug"),
			},
			{
				Name:     "public",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Public"),
			},
			{
				Name:     "regions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Regions"),
			},
			{
				Name:     "min_disk_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinDiskSize"),
			},
			{
				Name:     "size_giga_bytes",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("SizeGigaBytes"),
			},
			{
				Name:     "created",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "error_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ErrorMessage"),
			},
		},
	}
}

func fetchImages(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {

	svc := meta.(*client.Client)

	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}

	done := false
	listFunc := func() error {
		data, resp, err := svc.Services.Images.List(ctx, opt)
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
