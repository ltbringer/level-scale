variable "ssh_private_key" {
  type      = string
  sensitive = true
}

variable "ssh_public_key" {
  type      = string
  sensitive = true
}

variable "hcloud_token" {
  sensitive = true
  default   = ""
}