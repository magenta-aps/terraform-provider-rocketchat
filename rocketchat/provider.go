package rocketchat

import (
	"fmt"
	"net/url"

	rocketmodels "github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	rocketsdk "github.com/RocketChat/Rocket.Chat.Go.SDK/rest"

	_ "unsafe"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type authInfo struct {
	token string
	id    string
}

type MyClient struct {
	*rocketsdk.Client

	auth *authInfo
}

func (c *MyClient) TokenLogin(user_id string, token string) error {
	if c.auth != nil {
		return nil
	}

	c.auth = &authInfo{id: user_id, token: token}
	// credentials.ID, credentials.Token = response.Data.UserID, response.Data.Token
	return nil
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROCKETCHAT_ENDPOINT", nil),
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if value == "" {
						errors = append(errors, fmt.Errorf("Endpoint must not be an empty string"))
					}

					return
				},
			},

			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROCKETCHAT_EMAIL", nil),
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					return
				},
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROCKETCHAT_PASSWORD", nil),
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					return
				},
			},

			"user_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROCKETCHAT_USER_ID", nil),
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					return
				},
			},

			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROCKETCHAT_TOKEN", nil),
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					return
				},
			},
		},

		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"rocketchat_channels": dataSourcesChannels(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	var email = d.Get("email").(string)
	var password = d.Get("password").(string)
	var endpoint = d.Get("endpoint").(string)
	var user_id = d.Get("user_id").(string)
	var token = d.Get("token").(string)

	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	raw_client := rocketsdk.NewClient(url, true)
	client := MyClient{raw_client, nil}

	if user_id != "" && token != "" {
		client.TokenLogin(user_id, token)
	} else {
		credentials := rocketmodels.UserCredentials{
			Email:    email,
			Password: password,
		}

		err = client.Login(&credentials)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}
