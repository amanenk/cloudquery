// Code generated by codegen; DO NOT EDIT.

package lambda

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FunctionEventInvokeConfigs() *schema.Table {
	return &schema.Table{
		Name:      "aws_lambda_function_event_invoke_configs",
		Resolver:  fetchLambdaFunctionEventInvokeConfigs,
		Multiplex: client.ServiceAccountRegionMultiplexer("lambda"),
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
				Name:     "function_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "destination_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DestinationConfig"),
			},
			{
				Name:     "last_modified",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModified"),
			},
			{
				Name:     "maximum_event_age_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaximumEventAgeInSeconds"),
			},
			{
				Name:     "maximum_retry_attempts",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaximumRetryAttempts"),
			},
		},
	}
}
