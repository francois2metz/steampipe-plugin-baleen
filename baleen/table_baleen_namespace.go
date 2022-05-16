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
			{Func: getWaf},
			{Func: getHeaders},
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
				Name:        "cdn_enabled",
				Hydrate:     getCache,
				Type:        proto.ColumnType_BOOL,
				Description: "Cache enabled.",
				Transform:   transform.FromField("Enabled"),
			},
			{
				Name:        "waf_enabled",
				Hydrate:     getWaf,
				Type:        proto.ColumnType_BOOL,
				Description: "Waf enabled.",
				Transform:   transform.FromField("Enabled"),
			},
			{
				Name:        "headers_deny_frame_options",
				Hydrate:     getHeaders,
				Type:        proto.ColumnType_BOOL,
				Description: "Add X-Frame-Options: deny http header.",
				Transform:   transform.FromField("DenyFrameOptions"),
			},
			{
				Name:        "headers_no_sniff_mime_type",
				Hydrate:     getHeaders,
				Type:        proto.ColumnType_BOOL,
				Description: "Add X-Content-Type-Options: nosniff http header.",
				Transform:   transform.FromField("NoSniffMimeType"),
			},
		},
	}
}

func listNamespace(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_owasp.listNamespace", "connection_error", err)
		return nil, err
	}
	account, err := client.GetAccount()
	if err != nil {
		plugin.Logger(ctx).Error("baleen_namespace.listNamespace", err)
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
		plugin.Logger(ctx).Error("baleen_owasp.getOrigin", "connection_error", err)
		return nil, err
	}

	origin, err := client.GetOrigin(namespace.ID)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_namespace.getOrigin", err)
		return nil, err
	}
	return origin, nil
}

func getCache(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getCache")
	namespace := h.Item.(baleen.Namespace)

	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_owasp.getCache", "connection_error", err)
		return nil, err
	}

	cache, err := client.GetCache(namespace.ID)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_namespace.getCache", err)
		return nil, err
	}
	return cache, nil
}

func getWaf(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getWaf")
	namespace := h.Item.(baleen.Namespace)

	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_owasp.getWaf", "connection_error", err)
		return nil, err
	}

	waf, err := client.GetWaf(namespace.ID)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_namespace.getWaf", err)
		return nil, err
	}
	return waf, nil
}

func getHeaders(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getHeaders")
	namespace := h.Item.(baleen.Namespace)

	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_owasp.getHeaders", "connection_error", err)
		return nil, err
	}

	headers, err := client.GetHeaders(namespace.ID)
	if err != nil {
		plugin.Logger(ctx).Error("baleen_namespace.getHeaders", err)
		return nil, err
	}
	return headers, nil
}
