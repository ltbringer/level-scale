terraform {
  required_version = ">= 1.3.0"
  required_providers {
    hcloud = {
      source  = "hetznercloud/hcloud"
      version = ">= 1.40.0"
    }
  }
  backend "remote" {
    organization = "level-scale"
    workspaces {
      name = "level-scale-dev-nodes"
    }
  }
}

provider "hcloud" {
  token = var.hcloud_token
}

variable "hcloud_token" {
  type      = string
  sensitive = true
}

resource "hcloud_network" "priv_net" {
  name     = "level-scale-private"
  ip_range = "10.0.0.0/16"
}

resource "hcloud_network_subnet" "priv_subnet" {
  network_id   = hcloud_network.priv_net.id
  type         = "cloud"
  network_zone = "eu-central"
  ip_range     = "10.0.0.0/24"
}

resource "hcloud_server" "arm_node" {
  count       = 2
  name        = "arm-node-${count.index + 1}"
  server_type = "cax11"
  image       = "ubuntu-22.04"
  location    = "nbg1"
  ssh_keys    = [hcloud_ssh_key.my_key.id]

  network {
    network_id = hcloud_network.priv_net.id
  }
}

resource "hcloud_ssh_key" "my_key" {
  name       = "my-local-key"
  public_key = file("~/.ssh/id_ed25519.pub")
}
