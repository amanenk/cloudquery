// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func InstanceSnapshots() *schema.Table {
	return &schema.Table{
		Name:      "aws_lightsail_instance_snapshots",
		Resolver:  fetchLightsailInstanceSnapshots,
		Multiplex: client.ServiceAccountRegionMultiplexer("lightsail"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "from_attached_disks",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FromAttachedDisks"),
			},
			{
				Name:     "from_blueprint_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FromBlueprintId"),
			},
			{
				Name:     "from_bundle_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FromBundleId"),
			},
			{
				Name:     "from_instance_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FromInstanceArn"),
			},
			{
				Name:     "from_instance_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FromInstanceName"),
			},
			{
				Name:     "is_from_auto_snapshot",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsFromAutoSnapshot"),
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "progress",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Progress"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "size_in_gb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SizeInGb"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "support_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SupportCode"),
			},
		},
	}
}
