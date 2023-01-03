// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	{{if not .SkipFetch}}
	"context"
	"google.golang.org/api/iterator"
	{{if .ProtobufImport}}
  pb "{{.ProtobufImport}}"
  {{end}}
	{{end}}
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	{{if not .SkipFetch}}
	{{range .MockImports}}
  "{{.}}"
  {{end}}
	{{end}}
  {{range .Imports}}
  "{{.}}"
  {{end}}
)

func {{.SubService | ToCamel}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

{{if not .SkipFetch}}
func fetch{{.SubService | ToCamel}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.{{.RequestStructName}}{
		{{if .RequestStructFields}}{{.RequestStructFields}}{{end}}
	}{{ if .ParentFilterFunc }}
	if {{.ParentFilterFunc}}(parent) {
		return nil
	}
	{{ end }}
	{{ if .ClientOptions }}
	clientOptions := c.ClientOptions
	clientOptions = append([]option.ClientOption{ {{ range .ClientOptions }}{{.}},{{end}} }, clientOptions...)
	gcpClient, err := {{ or .ServiceAPIOverride .Service }}.{{.NewFunctionName}}(ctx, clientOptions...)
	{{ else }}gcpClient, err := {{ or .ServiceAPIOverride .Service }}.{{.NewFunctionName}}(ctx, c.ClientOptions...){{ end }}
	if err != nil {
		return err
	}
  it := gcpClient.{{.ListFunctionName}}(ctx, req, c.CallOptions...)
	for {
    resp, err := it.Next()
    if err == iterator.Done {
            break
    }
    if err != nil {
      return err
    }
		{{if .OutputField}}
			res <- resp.{{.OutputField}}
		{{else}}
			res <- resp
		{{end}}
	}
	return nil
}
{{end}}
