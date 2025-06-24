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
When using the `aws_ami` data source in Terraform with `most_recent = true`, omitting the `owners` attribute or specific filters such as `owner-id`, `owner-alias`, or `image-id` can cause Terraform to select AMIs from unknown or unauthorized sources. This increases the risk of deploying instances from untrusted or potentially malicious AMIs, which may introduce security vulnerabilities or unexpected behavior in your infrastructure. To mitigate this risk, always specify an explicit owner or unique identifier when querying for the most recent AMI.

Example of unsecure configuration:

```
data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }
}
```

Example of secure configuration:

```
data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"] // Canonical

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ami)


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