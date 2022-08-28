// Code generated by codegen; DO NOT EDIT.

package monitoring

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func AlertPolicies() *schema.Table {
	return &schema.Table{
		Name:      "gcp_monitoring_alert_policies",
		Resolver:  fetchAlertPolicies,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "alert_strategy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AlertStrategy"),
			},
			{
				Name:     "combiner",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Combiner"),
			},
			{
				Name:     "conditions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Conditions"),
			},
			{
				Name:     "creation_record",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CreationRecord"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "documentation",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Documentation"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
			{
				Name:     "mutation_record",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MutationRecord"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "notification_channels",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("NotificationChannels"),
			},
			{
				Name:     "user_labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UserLabels"),
			},
			{
				Name:     "validity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Validity"),
			},
		},
	}
}

func fetchAlertPolicies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Monitoring.Projects.AlertPolicies.List("projects/" + c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		res <- output.AlertPolicies

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
