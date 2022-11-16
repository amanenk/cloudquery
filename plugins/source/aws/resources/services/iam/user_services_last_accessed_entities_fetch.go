package iam

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIamUserServicesLastAccessedEntities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(models.ServiceLastAccessedEntitiesWrapper)
	c := meta.(*client.Client)
	svc := c.Services().Iam
	return fetchDetailEntities(ctx, res, svc, p)
}
