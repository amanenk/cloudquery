// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"github.com/cloudquery/plugin-sdk/schema"
	{{- if .ImportClient }}
	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	{{- end }}
)

func {{.TableFuncName}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}