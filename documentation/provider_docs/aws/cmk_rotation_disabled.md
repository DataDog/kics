---
title: "CMK Rotation Disabled"
meta:
  name: "aws/cmk_rotation_disabled"
  id: "22fbfeac-7b5a-421a-8a27-7a2178bb910b"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `aws/cmk_rotation_disabled`

**Id:** `22fbfeac-7b5a-421a-8a27-7a2178bb910b`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Observability

## Description
Customer Master Keys (CMK) must have rotation enabled, which means the attribute 'enable_key_rotation' must be set to 'true' when the key is enabled.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_key#enable_key_rotation)

## Non-Compliant Code Examples
```terraform
resource "aws_kms_key" "positive2" {
  description         = "KMS key 2"
  is_enabled          = true
  enable_key_rotation = false
}

```

```terraform
resource "aws_kms_key" "positive3" {
  description              = "KMS key 3"
  is_enabled               = true
  customer_master_key_spec = "SYMMETRIC_DEFAULT"
  enable_key_rotation      = false
}

```

```terraform
resource "aws_kms_key" "positive4" {
  description              = "KMS key 4"
  customer_master_key_spec = "SYMMETRIC_DEFAULT"
  enable_key_rotation      = false
}

```

```terraform
resource "aws_kms_key" "positive5" {
  description              = "KMS key 5"
  customer_master_key_spec = "RSA_2048"
  enable_key_rotation      = true
}

```

```terraform
resource "aws_kms_key" "positive1" {
  description = "KMS key 1"
}

```

## Compliant Code Examples
```terraform
resource "aws_kms_key" "negative3" {
  description              = "KMS key 3"
  customer_master_key_spec = "RSA_2048"
}

```

```terraform
resource "aws_kms_key" "negative2" {
  description              = "KMS key 2"
  customer_master_key_spec = "RSA_4096"
}

```

```terraform
resource "aws_kms_key" "negative1" {
  description         = "KMS key 1"
  is_enabled          = true
  enable_key_rotation = true
}

```

```terraform
resource "aws_kms_key" "negative5" {
  description              = "KMS key 5"
  customer_master_key_spec = "SYMMETRIC_DEFAULT"
  enable_key_rotation      = true
}

```

```terraform
resource "aws_kms_key" "negative4" {
  description              = "KMS key 4"
  customer_master_key_spec = "RSA_3072"
}

```