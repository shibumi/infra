resource "github_app_installation_repository" "netflify" {
  installation_id = var.github_netlify_app_id
  repository      = var.github_netlify_repo_name
}