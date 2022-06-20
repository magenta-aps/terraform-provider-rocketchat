terraform {}

terraform {
  required_providers {
    rocketchat = {
      source  = "magenta-aps/rocketchat"
      version = "1.0.0"
    }
  }
  required_version = ">= 1.1.4"
}

provider "rocketchat" {
  endpoint = "http://localhost:3000"
  username = "admin"
  password = "admin"
}

data "rocketchat_channels" "all_channels" {}

output "debug" {
  value = data.rocketchat_channels.all_channels
}
