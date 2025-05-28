---
title: "EFS Without KMS"
meta:
  name: "aws/efs_without_kms"
  id: "25d251f3-f348-4f95-845c-1090e41a615c"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Encryption"
---

## Metadata
**Name:** `aws/efs_without_kms`

**Id:** `25d251f3-f348-4f95-845c-1090e41a615c`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Encryption

## Description
Amazon Elastic Filesystem should have filesystem encryption enabled using KMS CMK customer-managed keys instead of AWS managed-keys

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/efs_file_system#kms_key_id)

## Non-Compliant Code Examples
```terraform
resource "aws_efs_file_system" "positive1" {
  creation_token = "my-product"
  encrypted = true

  tags = {
    Name = "MyProduct"
  }
}
```

## Compliant Code Examples
```terraform
resource "aws_efs_file_system" "negative1" {
  creation_token = "my-product"
  encrypted = true
  kms_key_id = "1234abcd-12ab-34cd-56ef-1234567890ab"

  tags = {
    Name = "MyProduct"
  }
}
```