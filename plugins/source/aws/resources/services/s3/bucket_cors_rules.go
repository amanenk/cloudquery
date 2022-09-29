// Code generated by codegen; DO NOT EDIT.

package s3

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func BucketCorsRules() *schema.Table {
	return &schema.Table{
		Name:      "aws_s3_bucket_cors_rules",
		Resolver:  fetchS3BucketCorsRules,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "bucket_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "allowed_methods",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AllowedMethods"),
			},
			{
				Name:     "allowed_origins",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AllowedOrigins"),
			},
			{
				Name:     "allowed_headers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AllowedHeaders"),
			},
			{
				Name:     "expose_headers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ExposeHeaders"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "max_age_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxAgeSeconds"),
			},
		},
	}
}
