// Code generated by codegen; DO NOT EDIT.

package wafv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ManagedRuleGroups() *schema.Table {
	return &schema.Table{
		Name:                 "aws_wafv2_managed_rule_groups",
		Resolver:             fetchWafv2ManagedRuleGroups,
		PostResourceResolver: resolveDescribeManagedRuleGroup,
		Multiplex:            client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "scope",
				Type:     schema.TypeString,
				Resolver: client.ResolveWAFScope,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "vendor_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VendorName"),
			},
			{
				Name:     "versioning_supported",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("VersioningSupported"),
			},
		},
	}
}
