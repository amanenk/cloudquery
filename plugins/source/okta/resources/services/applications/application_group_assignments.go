// Code generated by codegen; DO NOT EDIT.

package applications

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApplicationGroupAssignments() *schema.Table {
	return &schema.Table{
		Name:     "okta_application_group_assignments",
		Resolver: fetchApplicationGroupAssignments,
		Columns: []schema.Column{
			{
				Name:     "app_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "last_updated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdated"),
			},
			{
				Name:     "priority",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Priority"),
			},
			{
				Name:     "profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Profile"),
			},
			{
				Name:     "_embedded",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Embedded"),
			},
			{
				Name:     "_links",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Links"),
			},
			{
				Name:     "additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdditionalProperties"),
			},
		},
	}
}