// Code generated by codegen; DO NOT EDIT.

package ses

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Templates() *schema.Table {
	return &schema.Table{
		Name:      "aws_ses_templates",
		Resolver:  fetchSesTemplates,
		Multiplex: client.ServiceAccountRegionMultiplexer("email"),
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
				Resolver: resolveSesTemplateArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "template_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TemplateName"),
			},
			{
				Name:     "html",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Html"),
			},
			{
				Name:     "subject",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subject"),
			},
			{
				Name:     "text",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Text"),
			},
			{
				Name:     "created_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTimestamp"),
			},
		},
	}
}
