---
title: "AMI Most Recent Without Owner or Filter"
meta:
  name: "aws/ami_owner_missing"
  id: "f317c3c2-58e5-4aa7-bb8e-3a7f6bcd274a"
  cloud_provider: "aws"
  severity: "HIGH"
  category: "Access Control"
---

## Metadata
**Name:** `aws/ami_owner_missing`

**Id:** `f317c3c2-58e5-4aa7-bb8e-3a7f6bcd274a`

**Cloud Provider:** aws

**Severity:** High

**Category:** Access Control

## Description
Avoid using the aws_ami data source with most_recent = true unless an 'owners' attribute or a specific filter like 'owner-id', 'owner-alias', or 'image-id' is set. Omitting these may result in unintended AMIs being selected.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ami)

## Non-Compliant Code Examples
```terraform
# non-compliant
data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }
}

```

## Compliant Code Examples
```terraform
data "aws_ami" "ubuntu" {
  most_recent = true
  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }
  filter {
    name   = "owner-id"
    values = ["099720109477"]
  }
}

data "aws_ami" "ubuntu" {
  most_recent = true
  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }
  filter {
    name   = "owner-alias"
    values = ["amazon"]
  }
}

data "aws_ami" "ubuntu" {
  most_recent = true
  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }
  owners = ["099720109477"] // Canonical
}

data "aws_ami" "ubuntu" {
  most_recent = true
  filter {
    name   = "image-id"
    values = ["ami-12345"]
  }
}

data "aws_ami" "ubuntu" {
  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-1234"]
  }
}

data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"] // Canonical
  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }
}

```