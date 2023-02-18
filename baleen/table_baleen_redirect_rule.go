package baleen

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBaleenRedirectRule() *plugin.Table {
	return &plugin.Table{
		Name:        "baleen_redirect_rule",
		Description: "A redirect rule allow users to access moved content while maintaining the natural referencing of the pages.",
		List: &plugin.ListConfig{
			Hydrate:    listRedirectRule,
			KeyColumns: plugin.SingleColumn("namespace"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("namespace"),
				Description: "Namespace of the redirect rule.",
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
				Name:        "type",
				Type:        proto.ColumnType_INT,
				Description: "Type of the redirection (301 or 302).",
			},
			{
				Name:        "with_query_string",
				Type:        proto.ColumnType_BOOL,
				Description: "Keep the query string.",
			},
		},
	}
}

func listRedirectRule(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_owasp.listRedirectRule", "connection_error", err)
		return nil, err
	}
	namespace := d.EqualsQuals["namespace"].GetStringValue()
	rules, err := client.GetUrlRules(namespace)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_owasp.listRedirectRule", err)
		return nil, err
	}
	for _, rule := range rules.RedirectRules {
		d.StreamListItem(ctx, rule)
	}
	return nil, nil
}
