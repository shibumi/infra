data "external" "hetzner_cloud_api_key" {
	program = ["${path.module}/fetch-key.sh"]
}

resource "hcloud_server" "kurisu" {
  name        = "kurisu"
  server_type = "cx11-ceph"
  location    = "fsn1"
  image       = "fedora-31"
  lifecycle {
    ignore_changes = [image]
  }
}

resource "hcloud_rdns" "kurisu" {
  server_id  = hcloud_server.kurisu.id
  ip_address = hcloud_server.kurisu.ipv4_address
  dns_ptr    = "kurisu.shibumi.dev"
}
