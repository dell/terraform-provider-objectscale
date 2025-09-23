terraform {
  required_providers {
    objectscale = {
      source = "registry.terraform.io/dell/objectscale"
    }
  }
}

provider "objectscale" {
  endpoint = "https://10.225.108.217:4443"
  username = "root"
  password = "Password123!"
  insecure = true
}

data "objectscale_namespace" "all" {
}

output "objectscale_namespace_all" {
  value = data.objectscale_namespace.all
}
