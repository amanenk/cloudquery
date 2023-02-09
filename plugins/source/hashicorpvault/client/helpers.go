package client

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
)

type Item struct {
	Path string
	Data map[string]interface{}
}

func ListHelper(ctx context.Context, meta schema.ClientMeta, path string, res chan<- any) error {
	c := meta.(*Client)
	list, err := c.Client.Logical().ListWithContext(ctx, path)
	if err != nil {
		return err
	}
	if list == nil {
		c.logger.Warn().Str("path", path).Msg("list result is nil")
	}
	for _, key := range list.Data["keys"].([]interface{}) {
		res <- path + "/" + key.(string)
	}
	return nil
}
