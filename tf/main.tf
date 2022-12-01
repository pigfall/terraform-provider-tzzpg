terraform {
  required_providers{
    tzzpg = {
      source = "registry.terraform.io/pigfall/tzzpg"
    }
  }
}

provider "tzzpg" {}


data "tzzpg_dog" "dog"{}
