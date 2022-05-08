package baleen

import (
	"context"

	baleen "github.com/francois2metz/steampipe-plugin-baleen/baleen/client"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableBaleenNamespace() *plugin.Table {
	return &plugin.Table{
		Name:        "baleen_namespace",
		Description: "A namespace is a proxy with a configured origin and specific rules.",
		List: &plugin.ListConfig{
			Hydrate: listNamespace,
		},
		HydrateDependencies: []plugin.HydrateDependencies{
			{Func: getOrigin},
			{Func: getCache},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Unique ID of the namespace.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the namespace.",
			},
			{
				Name:        "url",
				Hydrate:     getOrigin,
				Type:        proto.ColumnType_STRING,
				Description: "URL of the origin.",
			},
			{
				Name:        "custom_404_page",
				Hydrate:     getOrigin,
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("ErrorPages.Custom404Page"),
				Description: "Use custom 404 page.",
			},
			{
				Name:        "custom_50x_page",
				Hydrate:     getOrigin,
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("ErrorPages.Custom50xPage"),
				Description: "Use custom 50x page.",
			},
			{
				Name:        "cache",
				Hydrate:     getCache,
				Type:        proto.ColumnType_BOOL,
				Description: "Cache enabled.",
				Transform:   transform.FromField("Enabled"),
			},
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

func getOrigin(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getOrigin")
	namespace := h.Item.(baleen.Namespace)

	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	origin, err := client.GetOrigin(namespace.ID)
	if err != nil {
		return nil, err
	}
	return origin, nil
}

func getCache(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getCache")
	namespace := h.Item.(baleen.Namespace)

	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	cache, err := client.GetCache(namespace.ID)
	if err != nil {
		return nil, err
	}
	return cache, nil
}
