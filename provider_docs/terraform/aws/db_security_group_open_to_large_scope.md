---
title: "DB Security Group Open To Large Scope"
meta:
  name: "terraform/db_security_group_open_to_large_scope"
  id: "4f615f3e-fb9c-4fad-8b70-2e9f781806ce"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/db_security_group_open_to_large_scope`
**Id:** `4f615f3e-fb9c-4fad-8b70-2e9f781806ce`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Networking and Firewall
## Description
This check ensures that DB Security Groups aren't configured with overly broad CIDR ranges that could expose your database to unnecessary network access. When a CIDR block with more than 256 hosts (such as /24 or lower) is configured in the ingress rules, it increases the attack surface and potential for unauthorized access to your database instances.

In the insecure example below, the security group allows access from a /24 CIDR block (256 hosts):
```
resource "aws_db_security_group" "positive1" {
  name = "rds_sg"

  ingress {
    cidr = "10.0.0.0/24"
  }
}
```

To remediate this issue, restrict access to a smaller CIDR range with fewer hosts, such as /25 (128 hosts) or more restrictive:
```
resource "aws_db_security_group" "negative1" {
  name = "rds_sg"

  ingress {
    cidr = "10.0.0.0/25"
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_security_group)

## Non-Compliant Code Examples
```aws
resource "aws_db_security_group" "positive1" {
  name = "rds_sg"

  ingress {
    cidr = "10.0.0.0/24"
  }
}
```

## Compliant Code Examples
```aws
resource "aws_db_security_group" "negative1" {
  name = "rds_sg"

  ingress {
    cidr = "10.0.0.0/25"
  }
}
```