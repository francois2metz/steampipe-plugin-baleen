package baleen

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableBaleenNamespace() *plugin.Table {
	return &plugin.Table{
		Name:        "baleen_namespace",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listNamespace,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique id of the namespace."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the namespace."},
		},
	}
}

func listNamespace(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	account, err := client.GetAccount()
	if err != nil {
		return nil, err
	}
	for _, namespace := range account.Namespaces {
		d.StreamListItem(ctx, namespace)
	}
	return nil, nil
}
