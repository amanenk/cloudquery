// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Alarms() *schema.Table {
	return &schema.Table{
		Name:      "aws_lightsail_alarms",
		Resolver:  fetchLightsailAlarms,
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
				Name:     "comparison_operator",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ComparisonOperator"),
			},
			{
				Name:     "contact_protocols",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ContactProtocols"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "datapoints_to_alarm",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DatapointsToAlarm"),
			},
			{
				Name:     "evaluation_periods",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("EvaluationPeriods"),
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "metric_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricName"),
			},
			{
				Name:     "monitored_resource_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MonitoredResourceInfo"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "notification_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("NotificationEnabled"),
			},
			{
				Name:     "notification_triggers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("NotificationTriggers"),
			},
			{
				Name:     "period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Period"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "statistic",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Statistic"),
			},
			{
				Name:     "support_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SupportCode"),
			},
			{
				Name:     "threshold",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("Threshold"),
			},
			{
				Name:     "treat_missing_data",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TreatMissingData"),
			},
			{
				Name:     "unit",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Unit"),
			},
		},
	}
}
