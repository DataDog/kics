---
title: "AMI Shared With Multiple Accounts"
meta:
  name: "aws/ami_shared_with_multiple_accounts"
  id: "ba4e0031-3e9d-4d7d-b0d6-bd8f003f8698"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/ami_shared_with_multiple_accounts`

**Id:** `ba4e0031-3e9d-4d7d-b0d6-bd8f003f8698`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Limits access to AWS AMIs by checking if more than one account is using the same image

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ami_launch_permission)

## Non-Compliant Code Examples
```terraform
resource "aws_ami_launch_permission" "positive1" {

  image_id   = "ami-1235678"
  account_id = "12345600012"

}


resource "aws_ami_launch_permission" "positive2" {

  image_id   = "ami-1235678"
  account_id = "123456789012"

}
```

## Compliant Code Examples
```terraform
resource "aws_ami_launch_permission" "negative1" {
  image_id   = "ami-12345678"
  account_id = "123456789012"
}


resource "aws_ami_launch_permission" "example" {
  image_id   = "ami-12345680"
  account_id = "12345672"
}

```