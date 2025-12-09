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
}

resource "objectscale_iam_role" "example" {
  name      = "example-role"
  namespace = "ns1"
  description = "An example role updated"
  permissions_boundary_arn = "urn:ecs:iam:::policy/ECSS3FullAccess"
  max_session_duration = 4000
  assume_role_policy_document = jsonencode({
    Version = "2012-11-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          AWS = [
            "urn:ecs:iam::ns1:user/sample_user_1"
          ]
        }
        Action = "sts:AssumeRole"
      }
    ]
  })
}