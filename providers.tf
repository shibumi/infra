terraform {
  required_version = ">= 1.0.0"
  required_providers {
    hcloud = {
      source  = "hetznercloud/hcloud"
      version = "~> 1.26.0"
    }
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 2.0"
    }
    netlify = {
      source  = "AegirHealth/netlify"
      version = "0.6.12"
    }
    github = {
      source  = "integrations/github"
      version = "~> 4.0"
    }
  }
}

provider "hcloud" {
  token = var.hetzner_cloud_api_key
}

provider "cloudflare" {
  email   = var.cloudflare_email
  api_key = var.cloudflare_api_key
}

provider "netlify" {
  token = var.netlify_api_key
}

provider "github" {
  owner = var.github_organization
  token = var.github_api_key
}