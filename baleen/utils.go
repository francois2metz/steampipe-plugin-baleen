package baleen

import (
	"context"
	"errors"
	"os"

	baleen "github.com/francois2metz/steampipe-plugin-baleen/baleen/client"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*baleen.Client, error) {
	// get baleen client from cache
	cacheKey := "baleen"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*baleen.Client), nil
	}

	baleenConfig := GetConfig(d.Connection)

	if &baleenConfig == nil {
		return nil, errors.New("You must have a baleen config file")
	}

	token := os.Getenv("BALEEN_TOKEN")

	if baleenConfig.Token != nil {
		token = *baleenConfig.Token
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	c := baleen.New(
		baleen.WithToken(token),
	)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, c)

	return c, nil
}
