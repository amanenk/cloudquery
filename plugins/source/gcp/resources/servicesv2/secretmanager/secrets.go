// Code generated by codegen; DO NOT EDIT.

package secretmanager

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func Secrets() *schema.Table {
	return &schema.Table{
		Name:      "gcp_secretmanager_secrets",
		Resolver:  fetchSecrets,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "create_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "expire_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExpireTime"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "replication",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Replication"),
			},
			{
				Name:     "rotation",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Rotation"),
			},
			{
				Name:     "topics",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Topics"),
			},
			{
				Name:     "ttl",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Ttl"),
			},
			{
				Name:     "version_aliases",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VersionAliases"),
			},
		},
	}
}

func fetchSecrets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.SecretManager.Projects.Secrets.List("projects/" + c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		res <- output.Secrets

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
