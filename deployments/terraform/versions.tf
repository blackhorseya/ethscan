terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "blackhorseya"

    workspaces {
      name = "ethscan"
    }
  }

  required_providers {
    cloudflare = {
      source = "cloudflare/cloudflare"
    }
  }
}
