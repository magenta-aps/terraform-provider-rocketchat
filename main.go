package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/magenta-aps/terraform-provider-rocketchat/rocketchat"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: rocketchat.Provider,
	})
}
