// Code generated by codegen; DO NOT EDIT.

package guardduty

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Detectors() *schema.Table {
	return &schema.Table{
		Name:                "aws_guardduty_detectors",
		Description:         `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetDetector.html`,
		Resolver:            fetchGuarddutyDetectors,
		PreResourceResolver: getDetector,
		Multiplex:           client.ServiceAccountRegionMultiplexer("guardduty"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGuarddutyARN(),
			},
			{
				Name: "id",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceRole"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "data_sources",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DataSources"),
			},
			{
				Name:     "finding_publishing_frequency",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FindingPublishingFrequency"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},

		Relations: []*schema.Table{
			DetectorMembers(),
		},
	}
}
