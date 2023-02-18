package baleen

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-baleen",
		DefaultTransform: transform.FromGo().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"baleen_custom_static_rule": tableBaleenCustomStaticRule(),
			"baleen_namespace":          tableBaleenNamespace(),
			"baleen_owasp":              tableBaleenOwasp(),
			"baleen_redirect_rule":      tableBaleenRedirectRule(),
			"baleen_rewrite_rule":       tableBaleenRewriteRule(),
		},
	}
	return p
}
