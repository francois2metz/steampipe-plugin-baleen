package baleen

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableBaleenOwasp() *plugin.Table {
	return &plugin.Table{
		Name:        "baleen_owasp",
		Description: "List Open Web Application Security Project rules.",
		List: &plugin.ListConfig{
			Hydrate:    listOwasp,
			KeyColumns: plugin.SingleColumn("namespace"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("namespace"),
				Description: "Namespace of the owasp rule.",
			},

			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the OWASP rule.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the rule.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the rule.",
			},
			{
				Name:        "group",
				Type:        proto.ColumnType_STRING,
				Description: "The group of the rule.",
			},
			{
				Name:        "enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Status of the rule.",
			},
		},
	}
}

type CrsThematicStatus struct {
	ID          string
	Name        string
	Description string
	Group       string
	Enabled     bool
}

func listOwasp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_owasp.listOwasp", "connection_error", err)
		return nil, err
	}
	namespace := d.KeyColumnQuals["namespace"].GetStringValue()
	thematics, err := client.GetCrsThematics(namespace)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_owasp.listOwasp.GetCrsThematics", err)
		return nil, err
	}
	waf, err := client.GetWaf(namespace)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_owasp.listOwasp.GetWaf", err)
		return nil, err
	}

	thematicStatus := map[string]bool{}
	for _, thematic := range waf.CrsThematics {
		thematicStatus[thematic.ID] = thematic.Enabled
	}
	for _, thematic := range thematics {
		enabled := false
		if v, ok := thematicStatus[thematic.ID]; ok {
			enabled = v
		}
		d.StreamListItem(ctx, CrsThematicStatus{
			ID:          thematic.ID,
			Name:        thematic.Name,
			Description: thematic.Description,
			Group:       thematic.Group,
			Enabled:     enabled,
		})
	}
	return nil, nil
}
