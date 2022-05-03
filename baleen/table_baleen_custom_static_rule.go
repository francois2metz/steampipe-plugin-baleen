package baleen

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableBaleenCustomStaticRule() *plugin.Table {
	return &plugin.Table{
		Name:        "baleen_custom_static_rule",
		Description: "A static rule allow to classify traffic with some conditions.",
		List: &plugin.ListConfig{
			Hydrate:    listCustomStaticRule,
			KeyColumns: plugin.SingleColumn("namespace"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("namespace"),
				Description: "Namespace of the rule.",
			},

			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Unique ID of the rule.",
			},
			{
				Name:        "tracking_id",
				Type:        proto.ColumnType_STRING,
				Description: "Tracking ID of the rule.",
			},
			{
				Name:        "category",
				Type:        proto.ColumnType_STRING,
				Description: "Category of the rule.",
			},
			{
				Name:        "enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Is activated.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "Description of the rule.",
			},
			{
				Name:        "conditions",
				Type:        proto.ColumnType_JSON,
				Description: "Conditions of the rule.",
			},
			{
				Name:        "labels",
				Type:        proto.ColumnType_JSON,
				Description: "Labels of the rule.",
			},
		},
	}
}

func listCustomStaticRule(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	namespace := d.KeyColumnQuals["namespace"].GetStringValue()
	rules, err := client.GetCustomStaticRules(namespace)
	if err != nil {
		return nil, err
	}
	for _, rule := range rules {
		d.StreamListItem(ctx, rule)
	}
	return nil, nil
}
