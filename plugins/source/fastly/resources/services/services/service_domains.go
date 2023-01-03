// Code generated by codegen; DO NOT EDIT.

package services

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func ServiceDomains() *schema.Table {
	return &schema.Table{
		Name:        "fastly_service_domains",
		Description: `https://developer.fastly.com/reference/api/services/domain/`,
		Resolver:    fetchServiceDomains,
		Columns: []schema.Column{
			{
				Name:     "comment",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Comment"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "deleted_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeletedAt"),
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
				Name:     "service_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ServiceVersion"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},
	}
}
