resource "netlify_deploy_key" "shibumi_dev_key" {}

resource "netlify_site" "shibumi_dev" {
  name = "shibumi.dev"
  repo {
    repo_path     = "shibumi/shibumi.dev"
    provider      = "github"
    deploy_key_id = netlify_deploy_key.key.id
  }
}