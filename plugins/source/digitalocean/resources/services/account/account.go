// Code generated by codegen; DO NOT EDIT.

package account

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func Account() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_account",
		Resolver: fetchAccount,
		Columns: []schema.Column{
			{
				Name:     "droplet_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DropletLimit"),
			},
			{
				Name:     "floating_ip_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("FloatingIPLimit"),
			},
			{
				Name:     "reserved_ip_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ReservedIPLimit"),
			},
			{
				Name:     "volume_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("VolumeLimit"),
			},
			{
				Name:     "email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Email"),
			},
			{
				Name:     "uuid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UUID"),
			},
			{
				Name:     "email_verified",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EmailVerified"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusMessage"),
			},
			{
				Name:     "team",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Team"),
			},
		},
	}
}

func fetchAccount(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	getFunc := func() error {
		response, _, err := svc.DoClient.Account.Get(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- *response
		return nil
	}

	err := client.ThrottleWrapper(ctx, svc, getFunc)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
