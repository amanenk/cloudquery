// Code generated by codegen; DO NOT EDIT.

package billing

import (
	"context"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/billing/apiv1/billingpb"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/billing/apiv1"
)

func BillingAccounts() *schema.Table {
	return &schema.Table{
		Name:        "gcp_billing_billing_accounts",
		Description: `https://cloud.google.com/billing/docs/reference/rest/v1/billingAccounts#BillingAccount`,
		Resolver:    fetchBillingAccounts,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudbilling.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
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
				Name:     "open",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Open"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "master_billing_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterBillingAccount"),
			},
		},

		Relations: []*schema.Table{
			Budgets(),
		},
	}
}

func fetchBillingAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListBillingAccountsRequest{}
	gcpClient, err := billing.NewCloudBillingClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListBillingAccounts(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}
