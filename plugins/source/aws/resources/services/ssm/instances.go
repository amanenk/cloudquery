// Code generated by codegen; DO NOT EDIT.

package ssm

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:      "aws_ssm_instances",
		Resolver:  fetchSsmInstances,
		Multiplex: client.ServiceAccountRegionMultiplexer("ssm"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveInstanceARN,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "activation_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActivationId"),
			},
			{
				Name:     "agent_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AgentVersion"),
			},
			{
				Name:     "association_overview",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AssociationOverview"),
			},
			{
				Name:     "association_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssociationStatus"),
			},
			{
				Name:     "computer_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ComputerName"),
			},
			{
				Name:     "ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IPAddress"),
			},
			{
				Name:     "iam_role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamRole"),
			},
			{
				Name:     "instance_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceId"),
			},
			{
				Name:     "is_latest_version",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsLatestVersion"),
			},
			{
				Name:     "last_association_execution_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastAssociationExecutionDate"),
			},
			{
				Name:     "last_ping_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastPingDateTime"),
			},
			{
				Name:     "last_successful_association_execution_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastSuccessfulAssociationExecutionDate"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "ping_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PingStatus"),
			},
			{
				Name:     "platform_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlatformName"),
			},
			{
				Name:     "platform_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlatformType"),
			},
			{
				Name:     "platform_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlatformVersion"),
			},
			{
				Name:     "registration_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("RegistrationDate"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "source_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceId"),
			},
			{
				Name:     "source_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceType"),
			},
		},

		Relations: []*schema.Table{
			InstanceComplianceItems(),
		},
	}
}
