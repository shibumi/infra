variable "cloudflare_email" {}
variable "cloudflare_api_key" {}
variable "hetzner_cloud_api_key" {}
variable "netlify_api_key" {}
variable "github_api_key" {}
variable "github_netlify_app_id" {
  description = "The netlify app ID"
  default     = "9175125"
  type        = string
}
variable "github_netlify_repo_name" {
  description = "The name of the github repository for the netlify blog"
  default     = "shibumi.dev"
  type        = string
}
variable "github_organization" {
  description = "The github username or organization"
  default     = "shibumi"
  type        = string
}