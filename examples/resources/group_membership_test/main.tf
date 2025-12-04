terraform {
  required_providers {
    objectscale = {
      source = "registry.terraform.io/dell/objectscale"
    }
  }
}

provider "objectscale" {
  username = "admin1"
  password = "ChangeMe"
  endpoint = "https://10.247.101.251:4443"
  insecure = true
  timeout  = 30
}

resource "objectscale_iam_group" "example" {
  name      = "example-group"
  namespace = "ns1"
}

resource "objectscale_iam_group_membership" "example_membership" {
  name      = objectscale_iam_group.example.name
  namespace = objectscale_iam_group.example.namespace
  user      = "test-user"
}