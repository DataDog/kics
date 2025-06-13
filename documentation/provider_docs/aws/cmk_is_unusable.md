---
title: "CMK Is Unusable"
meta:
  name: "aws/cmk_is_unusable"
  id: "7350fa23-dcf7-4938-916d-6a60b0c73b50"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Availability"
---

## Metadata
**Name:** `aws/cmk_is_unusable`

**Id:** `7350fa23-dcf7-4938-916d-6a60b0c73b50`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Availability

## Description
AWS Key Management Service (KMS) must only possess usable Customer Master Keys (CMK), which means the CMKs must have the attribute 'is_enabled' set to true

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_key#is_enabled)

## Non-Compliant Code Examples
```terraform
resource "aws_kms_key" "a" {
  description             = "KMS key 1"
  is_enabled = false
}

```

## Compliant Code Examples
```terraform
resource "aws_kms_key" "a3" {
  description             = "KMS key 1"
  is_enabled = true
}

```