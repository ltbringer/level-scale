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
      name = "level-scale-dev-kubernetes"
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

variable "ssh_public_key" {
  type      = string
  sensitive = false
}

module "hetzner-cluster" {
  source = "git::https://github.com/poseidon/typhoon//hetzner-cluster?ref=v1.30.1"

  cluster_name = "level-scale-dev"
  region       = "nbg1"
  os_channel   = "flatcar-stable"

  controller_count = 1
  worker_count     = 2
  ssh_fingerprints = [hcloud_ssh_key.my_key.fingerprint]

  controller_type  = "cpx31"
  worker_type      = "cpx31"
}

resource "hcloud_ssh_key" "my_key" {
  name       = "my-local-key"
  public_key = var.ssh_public_key
}
