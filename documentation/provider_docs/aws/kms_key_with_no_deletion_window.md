---
title: "KMS Key With No Deletion Window"
meta:
  name: "aws/kms_key_with_no_deletion_window"
  id: "0b530315-0ea4-497f-b34c-4ff86268f59d"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `aws/kms_key_with_no_deletion_window`

**Id:** `0b530315-0ea4-497f-b34c-4ff86268f59d`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Observability

## Description
AWS KMS Key should have a valid deletion window

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_key)

## Non-Compliant Code Examples
```terraform
resource "aws_kms_key" "positive1" {
  description             = "KMS key 1"
  
  is_enabled = true

  enable_key_rotation = true

}


resource "aws_kms_key" "positive2" {
  description             = "KMS key 1"
  
  is_enabled = true

  enable_key_rotation = true

  deletion_window_in_days = 31
}

```

## Compliant Code Examples
```terraform
resource "aws_kms_key" "negative1" {
  description             = "KMS key 1"
  
  is_enabled = true

  enable_key_rotation = true

  deletion_window_in_days = 10
}
```