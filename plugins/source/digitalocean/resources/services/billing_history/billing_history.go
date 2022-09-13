// Code generated by codegen; DO NOT EDIT.

package billing_history

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/digitalocean/godo"
)

func BillingHistory() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_billing_history",
		Resolver: fetchBillingHistory,
		Columns: []schema.Column{
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "amount",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Amount"),
			},
			{
				Name:     "invoice_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InvoiceID"),
			},
			{
				Name:     "invoice_uuid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InvoiceUUID"),
			},
			{
				Name:     "date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Date"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchBillingHistory(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {

	svc := meta.(*client.Client)

	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}

	done := false
	listFunc := func() error {
		data, resp, err := svc.Services.BillingHistory.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's data to our result channel
		res <- data.BillingHistory
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
