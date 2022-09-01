package main

import (
	"github.com/francois2metz/steampipe-plugin-baleen/baleen"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: baleen.Plugin})
}
