terraform {
  backend "remote" {
    organization = "level-scale"

    workspaces {
      name = "level-scale-dev-cluster"
    }
  }
}