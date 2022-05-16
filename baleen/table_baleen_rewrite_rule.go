package baleen

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableBaleenRewriteRule() *plugin.Table {
	return &plugin.Table{
		Name:        "baleen_rewrite_rule",
		Description: "A rewrite rule allow users to continue to access moved content seamlessly.",
		List: &plugin.ListConfig{
			Hydrate:    listRewriteRule,
			KeyColumns: plugin.SingleColumn("namespace"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("namespace"),
				Description: "Namespace of the rewrite rule.",
			},

			{
				Name:        "source",
				Type:        proto.ColumnType_STRING,
				Description: "Source path.",
			},
			{
				Name:        "destination",
				Type:        proto.ColumnType_STRING,
				Description: "Destination URL.",
			},
			{
				Name:        "with_query_string",
				Type:        proto.ColumnType_BOOL,
				Description: "Keep the query string.",
			},
		},
	}
}

func listRewriteRule(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_owasp.listRewriteRule", "connection_error", err)
		return nil, err
	}
	namespace := d.KeyColumnQuals["namespace"].GetStringValue()
	rules, err := client.GetUrlRules(namespace)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_owasp.listRewriteRule", err)
		return nil, err
	}
	for _, rule := range rules.RewriteRules {
		d.StreamListItem(ctx, rule)
	}
	return nil, nil
}
