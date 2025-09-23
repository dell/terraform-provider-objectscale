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

resource "objectscale_namespace" "example" {
  name = "luis_namespace"
  default_data_services_vpool = "urn:storageos:ReplicationGroupInfo:0e953ad1-94a5-4eb1-825a-d58d29e85434:global"
  retention_classes = {
    retention_class = [{
      name = "r1"
      period = 1
    }]
  }
  user_mapping = [{
    attributes = [{
      key = "key1"
      value = ["value1"]
    }]
    domain =  "domain"
    groups = ["group"]
  }]
}
