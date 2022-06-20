package rocketchat

import (
	"context"

	// rocketsdk "github.com/RocketChat/Rocket.Chat.Go.SDK/rest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcesChannels() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcesReadChannels,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "The id of the channel.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "The name of the channel.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"full_name": {
							Description: "The full name of the channel.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"type": {
							Description: "The channel type.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"messages": {
							Description: "Total number of messages in the channel.",
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"read_only": {
							Description: "Whether the channel is read only.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"sys_mes": {
							Description: "Whether the channel shows system messages.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"default": {
							Description: "Whether the channel is the default channel.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"broadcast": {
							Description: "Whether the channel is a broadcast channel.",
							Type:        schema.TypeBool,
							Computed:    true,
						},

						"creation_time": {
							Description: "When the channel was created.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"update_time": {
							Description: "When the channel was last updated.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourcesReadChannels(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(MyClient)
	channels, err := client.GetPublicChannels()
	if err != nil {
		return diag.FromErr(err)
	}
	values := []map[string]interface{}{}
	id := ""
	for _, channel := range channels.Channels {
		v := map[string]interface{}{
			"id":        channel.ID,
			"name":      channel.Name,
			"full_name": channel.Fname,
			"type":      channel.Type,
			"messages":  channel.Msgs,

			"read_only": channel.ReadOnly,
			"sys_mes":   channel.SysMes,
			"default":   channel.Default,
			"broadcast": channel.Broadcast,

			"creation_time": channel.Timestamp.Unix(),
			"update_time":   channel.UpdatedAt.Unix(),
		}
		values = append(values, v)
		id = id + channel.ID
	}

	d.Set("name", values)
	d.SetId(id)

	return diags
}
