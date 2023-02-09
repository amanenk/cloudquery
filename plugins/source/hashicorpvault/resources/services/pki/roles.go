package pki

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/hashicorpvault/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Roles() *schema.Table {
	return &schema.Table{
		Name:                "hashicorpvault_pki_roles",
		PreResourceResolver: resolvePkiRole,
		Description:         `https://developer.hashicorp.com/vault/api-docs/secret/pki#sample-response-1`,
		Resolver:            fetchPkiRoles,
		Transform:           transformers.TransformWithStruct(&client.Item{}, transformers.WithPrimaryKeys("Path")),
		Columns: []schema.Column{
			{
				Name:     "vault_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveVaultName,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func resolvePkiRole(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client)
	path := resource.Item.(string)
	item, err := svc.Client.Logical().ReadWithContext(ctx, path)
	if err != nil {
		return err
	}

	resource.SetItem(client.Item{
		Path: path,
		Data: item.Data,
	})

	return nil
}

func fetchPkiRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	for _, path := range c.Paths.Pki {
		err := client.ListHelper(ctx, meta, path+"/roles", res)
		if err != nil {
			return err
		}
	}
	return nil
}
